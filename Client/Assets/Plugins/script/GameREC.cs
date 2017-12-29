using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
using System;

namespace game {
	public enum RECResultCode
	{
		kRECInitSuccess = 0,/**< enum value is callback of succeeding in initing sdk . */
		kRECInitFail,/**< enum  value is callback of failing to init sdk. */
		kRECStartRecording,/**< enum  value is callback of starting to record. */
		kRECStopRecording,/**< enum  value is callback of stoping to record. */
		kRECPauseRecording,/**< enum  value is callback of pausing to record. */
		kRECResumeRecording,/**< enum  value is callback of resuming to record. */
		kRECEnterSDKPage,/**< enum  value is callback of failing to init sdk. */
		kRECQuitSDKPage,/**< enum  value is callback of entering SDK`s page. */
		kRECShareSuccess,/**< enum  value is callback of  quiting SDK`s page. */
		kRECShareFail,/**< enum  value is callback of failing to share. */
		kRECExtension = 90000 /**< enum value is  extension code . */
	};
	public class GameREC
	{

		private static GameREC _instance;
		
		public static GameREC getInstance() {
			if( null == _instance ) {
				_instance = new GameREC();
			}
			return _instance;
		}

		/**
     	@brief Start to record
    	 */
		
		public  void startRecording()
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			GameREC_nativeStartRecording ();
#else
			Debug.Log("This platform does not support!");
#endif
		}
		
		/**
     	@brief Stop a session.
    	 @warning This interface only worked on android
    	 */
		
		public  void stopRecording()
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			GameREC_nativeStopRecording ();
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
    	@brief share video
   		@param info The info of share, contains key:
   		 @warning Look at the manual of plugins.

    	*/
		
		public  void share(Dictionary<string,string> shareInfo)
		{
			
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			string info = GameUtil.dictionaryToString (shareInfo);
			Debug.Log("share   " + info);
			GameREC_nativeShare (info );
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
     	@brief Check function the plugin support or not
     	*/
		
		public  bool  isFunctionSupported (string functionName)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			return GameREC_nativeIsFunctionSupported (functionName);
#else
			Debug.Log("This platform does not support!");
			return false;
#endif
		}

		/**
		 * set debugmode for plugin
		 * 
		 */
		[Obsolete("This interface is obsolete!",false)]
		public  void setDebugMode(bool bDebug)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			GameREC_nativeSetDebugMode (bDebug);
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
    	 @brief set pListener The callback object for share result
    	 @param the MonoBehaviour object
    	 @param the callback of function
    	 */

		public  void setListener(MonoBehaviour gameObject,string functionName)
		{
#if !UNITY_EDITOR && UNITY_ANDROID
			GameUtil.registerActionCallback (GameType.REC, gameObject, functionName);
#elif !UNITY_EDITOR && UNITY_IOS
			string gameObjectName = gameObject.gameObject.name;
			GameREC_nativeSetListener(gameObjectName,functionName);
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
		 * Get Plugin version
		 * 
		 * @return string
	 	*/
		public  string getPluginVersion()
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			StringBuilder version = new StringBuilder();
			version.Capacity = GameUtil.MAX_CAPACITY_NUM;
			GameREC_nativeGetPluginVersion (version);
			return version.ToString();
#else
			Debug.Log("This platform does not support!");
			return "";
#endif
		}

		
		/**
		 * Get SDK version
		 * 
		 * @return string
	 	*/
		public  string getSDKVersion()
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			StringBuilder version = new StringBuilder();
			version.Capacity = GameUtil.MAX_CAPACITY_NUM;
			GameREC_nativeGetSDKVersion (version);
			return version.ToString();
#else
			Debug.Log("This platform does not support!");
			return "";
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param GameParam param 
    	 *@return void
    	 */
		
		public  void callFuncWithParam(string functionName, GameParam param)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS)  
			List<GameParam> list = new List<GameParam> ();
			list.Add (param);
			GameREC_nativeCallFuncWithParam(functionName, list.ToArray(),list.Count);
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param List<GameParam param 
    	 *@return void
    	 */
		public  void callFuncWithParam(string functionName, List<GameParam> param = null)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			if (param == null) 
			{
				GameREC_nativeCallFuncWithParam (functionName, null, 0);
				
			} else {
				GameREC_nativeCallFuncWithParam (functionName, param.ToArray (), param.Count);
			}
#else
			Debug.Log("This platform does not support!");
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param GameParam param 
    	 *@return int
    	 */
		public  int callIntFuncWithParam(string functionName, GameParam param)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS)  
			List<GameParam> list = new List<GameParam> ();
			list.Add (param);
			return GameREC_nativeCallIntFuncWithParam(functionName, list.ToArray(),list.Count);
#else
			Debug.Log("This platform does not support!");
			return -1;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param List<GameParam param 
    	 *@return int
    	 */
		public  int  callIntFuncWithParam(string functionName, List<GameParam> param = null)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS)  
			if (param == null)
			{
				return GameREC_nativeCallIntFuncWithParam (functionName, null, 0);
				
			} else {
				return GameREC_nativeCallIntFuncWithParam (functionName, param.ToArray (), param.Count);
			}
#else
			Debug.Log("This platform does not support!");
			return -1;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param GameParam param 
    	 *@return float
    	 */
		public  float callFloatFuncWithParam(string functionName, GameParam param)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			List<GameParam> list = new List<GameParam> ();
			list.Add (param);
			return GameREC_nativeCallFloatFuncWithParam(functionName, list.ToArray(),list.Count);
#else
			Debug.Log("This platform does not support!");
			return 0;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param List<GameParam param 
    	 *@return float
    	 */
		public  float callFloatFuncWithParam(string functionName, List<GameParam> param = null)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			if (param == null)
			{
				return GameREC_nativeCallFloatFuncWithParam (functionName, null, 0);
				
			} else {
				return GameREC_nativeCallFloatFuncWithParam (functionName, param.ToArray (), param.Count);
			}
#else
			Debug.Log("This platform does not support!");
			return 0;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param GameParam param 
    	 *@return bool
    	 */
		public  bool callBoolFuncWithParam(string functionName, GameParam param)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			List<GameParam> list = new List<GameParam> ();
			list.Add (param);
			return GameREC_nativeCallBoolFuncWithParam(functionName, list.ToArray(),list.Count);
#else
			Debug.Log("This platform does not support!");
			return false;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param List<GameParam param 
    	 *@return bool
    	 */
		public  bool callBoolFuncWithParam(string functionName, List<GameParam> param = null)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			if (param == null)
			{
				return GameREC_nativeCallBoolFuncWithParam (functionName, null, 0);
				
			} else {
				return GameREC_nativeCallBoolFuncWithParam (functionName, param.ToArray (), param.Count);
			}
#else
			Debug.Log("This platform does not support!");
			return false;
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param GameParam param 
    	 *@return string
    	 */
		public  string callStringFuncWithParam(string functionName, GameParam param)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			List<GameParam> list = new List<GameParam> ();
			list.Add (param);
			StringBuilder value = new StringBuilder();
			value.Capacity = GameUtil.MAX_CAPACITY_NUM;
			GameREC_nativeCallStringFuncWithParam(functionName, list.ToArray(),list.Count,value);
			return value.ToString ();
#else
			Debug.Log("This platform does not support!");
			return "";
#endif
		}

		/**
    	 *@brief methods for reflections
   	 	 *@param function name
   		 *@param List<GameParam param 
    	 *@return string
    	 */
		public  string callStringFuncWithParam(string functionName, List<GameParam> param = null)
		{
#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			StringBuilder value = new StringBuilder();
			value.Capacity = GameUtil.MAX_CAPACITY_NUM;
			if (param == null)
			{
				GameREC_nativeCallStringFuncWithParam (functionName, null, 0,value);
				
			} else {
				GameREC_nativeCallStringFuncWithParam (functionName, param.ToArray (), param.Count,value);
			}
			return value.ToString ();
#else
			Debug.Log("This platform does not support!");
			return "";
#endif
		}
		
		[DllImport(GameUtil.GAME_PLATFORM,CallingConvention=CallingConvention.Cdecl)]
		private static extern void GameREC_RegisterExternalCallDelegate(IntPtr functionPointer);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeStartRecording();
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeStopRecording();

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeShare(string info);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeSetListener(string gameName, string functionName);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern bool GameREC_nativeIsFunctionSupported(string functionName);
	
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeSetDebugMode(bool bDebug);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeGetPluginVersion(StringBuilder version);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeGetSDKVersion(StringBuilder version);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeCallFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern int GameREC_nativeCallIntFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern float GameREC_nativeCallFloatFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern bool GameREC_nativeCallBoolFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameREC_nativeCallStringFuncWithParam(string functionName, GameParam[] param,int count,StringBuilder value);
	}
	
}


