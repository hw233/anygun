-module(bin_reader).

-include("bin.hrl").

%%----------------------------------------------------------------------
%% Exported Functions
%%----------------------------------------------------------------------
-export([
	rd_type/5
	]).
%%----------------------------------------------------------------------
%% Func: rd_type/4
%% Args: 
%%		Type - data type.
%%		Array Max Size.
%% 		Field Max Size.
%%		Binary
%%		Mask - false, 1, 0.
%% Returns:
%%		{Value, RestBinary}  	
%%----------------------------------------------------------------------
%% array.
rd_type(T, ArrMax, StrMax, B, Mask) when ArrMax > 0 ->
	if
		Mask == 0 -> {[], B};
		true ->
			{Len, R} = rd_dyn_size(B),
			if
				Len =< ArrMax -> 
					rd_array_item(T, Len, StrMax, [], R)
			end
	end;
%% normal types.
rd_type(int64, 0, 0, B, Mask) -> 
	if
		Mask == 0 -> {0, B};
		true -> 
			<<V:64/little-signed-integer, R/binary>> = B, 
			{V, R}
	end;
rd_type(uint64, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:64/little-unsigned-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(double, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:64/little-float, R/binary>> = B,
			{V, R}
	end;
rd_type(float, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:32/little-float, R/binary>> = B,
			{V, R}
	end;
rd_type(int32, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:32/little-signed-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(uint32, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:32/little-unsigned-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(int16, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:16/little-signed-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(uint16, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:16/little-unsigned-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(int8, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:8/little-signed-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(uint8, 0, 0, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:8/little-unsigned-integer, R/binary>> = B,
			{V, R}
	end;
rd_type(bool, 0, 0, B, Mask) ->
	case Mask of
		false ->
			case B of
				<<0:8, R/binary>> -> {false, R};
				<<1:8, R/binary>> -> {true, R}
			end;
		0 -> {false, B};
		1 -> {true, B}
	end;
rd_type(string, 0, StrMax, B, Mask) ->
	if
		Mask == 0 -> {"", B};
		true ->
			{Len, R} = rd_dyn_size(B),
			if
				Len =< StrMax ->
					<<Str:Len/binary, R1/binary>> = R,
					{binary_to_list(Str), R1}
			end
	end;
%% enum.
rd_type(enum, 0, EnumMax, B, Mask) ->
	if
		Mask == 0 -> {0, B};
		true ->
			<<V:8/little-unsigned-integer, R/binary>> = B,
			if V < EnumMax -> {V, R} end
	end;
%% user type.
rd_type(UT, 0, 0, B, _) ->
	UT:deserialize(B).


%%----------------------------------------------------------------------
%% Local Functions
%%----------------------------------------------------------------------

rd_array_item(_, 0, _, L, B) ->
	{lists:reverse(L), B};
rd_array_item(T, N, StrMax, L, B) ->
	{V, R} = rd_type(T, 0, StrMax, B, false),
	rd_array_item(T, N-1, StrMax, [V|L], R).

%% dynamic size
rd_dyn_size(<<0:2, S:6, R/binary>>) ->
	{S, R};
rd_dyn_size(<<1:2, S:14/big-unsigned-integer, R/binary>>) ->
	{S, R};
rd_dyn_size(<<2:2, S:22/big-unsigned-integer, R/binary>>) ->
	{S, R};
rd_dyn_size(<<3:2, S:30/big-unsigned-integer, R/binary>>) ->
	{S, R}.