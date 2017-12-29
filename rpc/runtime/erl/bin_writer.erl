-module(bin_writer).

-include("bin.hrl").

%%----------------------------------------------------------------------
%% Exported Functions
%%----------------------------------------------------------------------
-export([
	wr_type/4
	]).

%%-----------------------------------------------------------------
%% Func: wr_type/4
%% Args: 
%%       Type - data type.
%%		 Array - is array.
%%       Value - data type value.
%%		 Mask - use mask.
%% Returns: 
%%       if Mask == true {Mask, Binary}
%%		 else Binary
%%-----------------------------------------------------------------
%% binary array.
wr_type(uint8, true, V, Mask) when is_binary(V) ->
	Len = iolist_size(V),
	case Mask of
		false ->
			iolist_to_binary([wr_dyn_size(Len),V]);
		true ->
			if
				Len == 0 ->
					{0, <<>>};
				true ->
					{1, iolist_to_binary([wr_dyn_size(Len),V])}
			end
	end;
%% array.
wr_type(T, true, V, Mask) ->
	Len = if V == undefined -> 0; is_list(V) -> length(V) end,
	if
		Len == 0 ->
			wr_0_b(wr_dyn_size(0), Mask);
		true ->
			DSize = wr_dyn_size(Len),
			Array = [wr_type(T, false, I, false) || I <- V],
			B = iolist_to_binary([DSize|Array]),
			wr_1_b(B, Mask)
	end;
%% normal types.
wr_type(int64, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:64/little-signed-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_INT64_MIN andalso V =< ?BIN_INT64_MAX ->
			wr_1_b(<<V:64/little-signed-integer>>, Mask)
	end;
wr_type(uint64, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:64/little-unsigned-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_UINT64_MIN andalso V =< ?BIN_UINT64_MAX ->
			wr_1_b(<<V:64/little-unsigned-integer>>, Mask)
	end;
wr_type(double, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 -> 
			wr_0_b(<<0:64/little-float>>, Mask);
		is_number(V) -> 
			wr_1_b(<<V:64/little-float>>, Mask)
	end;
wr_type(float, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 -> 
			wr_0_b(<<0:32/little-float>>, Mask);
		is_number(V) -> 
			wr_1_b(<<V:32/little-float>>, Mask)
	end;
wr_type(int32, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:32/little-signed-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_INT32_MIN andalso V =< ?BIN_INT32_MAX ->
			wr_1_b(<<V:32/little-signed-integer>>, Mask)
	end;
wr_type(uint32, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:32/little-unsigned-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_UINT32_MIN andalso V =< ?BIN_UINT32_MAX ->
			wr_1_b(<<V:32/little-unsigned-integer>>, Mask)
	end;
wr_type(int16, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:16/little-signed-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_INT16_MIN andalso V =< ?BIN_INT16_MAX ->
			wr_1_b(<<V:16/little-signed-integer>>, Mask)
	end;
wr_type(uint16, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:16/little-unsigned-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_UINT16_MIN andalso V =< ?BIN_UINT16_MAX ->
			wr_1_b(<<V:16/little-unsigned-integer>>, Mask)
	end;
wr_type(int8, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:8/little-signed-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_INT8_MIN andalso V =< ?BIN_INT8_MAX ->
			wr_1_b(<<V:8/little-signed-integer>>, Mask)
	end;
wr_type(uint8, false, V, Mask) ->
	if
		V == undefined orelse V =:= 0 ->
			wr_0_b(<<0:8/little-unsigned-integer>>, Mask);
		is_integer(V) andalso V >= ?BIN_UINT8_MIN andalso V =< ?BIN_UINT8_MAX ->
			wr_1_b(<<V:8/little-unsigned-integer>>, Mask)
	end;
wr_type(bool, false, V, Mask) ->
	if
		V == undefined orelse V == false -> 
			case Mask of
				false -> <<0:8>>;
				true -> {0, <<>>}
			end;
		V == true -> 
			case Mask of
				false -> <<1:8>>;
				true -> {1, <<>>}
			end
	end;
wr_type(string, false, V, Mask) ->
	Len = if V == undefined -> 0; is_list(V) -> length(V) end,
	if
		Len == 0 ->
			wr_0_b(wr_dyn_size(0), Mask);
		true ->
			wr_1_b(iolist_to_binary([wr_dyn_size(Len),V]), Mask)
	end;
%% enum
wr_type(enum, false, V, Mask) ->
	wr_type(uint8, false, V, Mask);
%% user type.
wr_type(UT, false, V, Mask) ->
	B = 
		if
			V == undefined -> UT:serialize(UT:default());
			is_tuple(V) -> UT:serialize(V)
		end,
	wr_1_b(B, Mask).

%%----------------------------------------------------------------------
%% Local Functions
%%----------------------------------------------------------------------
wr_0_b(B, Mask) ->
	case Mask of
		false -> B;
		true -> {0, <<>>}
	end.
wr_1_b(B, Mask) ->
	case Mask of
		false -> B;
		true -> {1, B}
	end.

%% dynamic size.
wr_dyn_size(S) when is_integer(S) andalso S >= 0 ->
	if
		S =< 16#3F ->
			<<0:2, S:6>>;
		S =< 16#3FFF ->
			<<1:2, S:14/big-unsigned-integer>>;
		S =< 16#3FFFFFFF ->
			<<2:2, S:22/big-unsigned-integer>>;
		true ->
			<<3:2, S:30/big-unsigned-integer>>
	end.