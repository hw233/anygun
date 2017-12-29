using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
using System;

namespace game {
	public enum ShareResultCode
	{
		kShareSuccess = 0,/**< enum value is callback of failing to sharing . */
		kShareFail,/**< enum value is callback of failing to share . */
		kShareCancel,/**< enum value is callback of canceling to share . */
		kShareNetworkError,/**< enum value is callback of network error . */
		kShareExtension = 10000 /**< enum value is  extension code . */
	};
	public class GameShare
	{

		private static GameShare _instance;
		
		public static GameShare getInstance() {
			if( null == _instance ) {
				_instance = new GameShare();
			}
			return _instance;
		}

		/**
    	@brief share information
   		@param info The info of share, contains key:
   		 @warning Look at the manual of plugins.

    	*/

		public  void share(Dictionary<string,string> shareInfo)
		{

#if !UNITY_EDITOR &&( UNITY_ANDROID ||  UNITY_IOS) 
			string info = GameUtil.dictionaryToString (shareInfo);
			Debug.Log("share   " + info);
			GameShare_nativeShare (info );
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
			return GameShare_nativeIsFunctionSupported (functionName);
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
			GameShare_nativeSetDebugMode (bDebug);
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
			GameUtil.registerActionCallback (GameType.Share, gameObject, functionName);
#elif !UNITY_EDITOR && UNITY_IOS
			string gameObjectName = gameObject.gameObject.name;
			GameShare_nativeSetListener(gameObjectName,functionName);
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
			GameShare_nativeGetPluginVersion (version);
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
			GameShare_nativeGetSDKVersion (version);
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
			GameShare_nativeCallFuncWithParam(functionName, list.ToArray(),list.Count);
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
				GameShare_nativeCallFuncWithParam (functionName, null, 0);
				
			} else {
				GameShare_nativeCallFuncWithParam (functionName, param.ToArray (), param.Count);
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
			return GameShare_nativeCallIntFuncWithParam(functionName, list.ToArray(),list.Count);
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
				return GameShare_nativeCallIntFuncWithParam (functionName, null, 0);
				
			} else {
				return GameShare_nativeCallIntFuncWithParam (functionName, param.ToArray (), param.Count);
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
			return GameShare_nativeCallFloatFuncWithParam(functionName, list.ToArray(),list.Count);
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
				return GameShare_nativeCallFloatFuncWithParam (functionName, null, 0);
				
			} else {
				return GameShare_nativeCallFloatFuncWithParam (functionName, param.ToArray (), param.Count);
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
			return GameShare_nativeCallBoolFuncWithParam(functionName, list.ToArray(),list.Count);
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
				return GameShare_nativeCallBoolFuncWithParam (functionName, null, 0);
				
			} else {
				return GameShare_nativeCallBoolFuncWithParam (functionName, param.ToArray (), param.Count);
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
			GameShare_nativeCallStringFuncWithParam(functionName, list.ToArray(),list.Count,value);
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
				GameShare_nativeCallStringFuncWithParam (functionName, null, 0,value);
				
			} else {
				GameShare_nativeCallStringFuncWithParam (functionName, param.ToArray (), param.Count,value);
			}
			return value.ToString ();
#else
			Debug.Log("This platform does not support!");
			return "";
#endif
		}
		
		[DllImport(GameUtil.GAME_PLATFORM,CallingConvention=CallingConvention.Cdecl)]
		private static extern void GameShare_RegisterExternalCallDelegate(IntPtr functionPointer);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeSetListener(string gameName, string functionName);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeShare(string info);

		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern bool GameShare_nativeIsFunctionSupported(string functionName);
	
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeSetDebugMode(bool bDebug);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeGetPluginVersion(StringBuilder version);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeGetSDKVersion(StringBuilder version);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeCallFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern int GameShare_nativeCallIntFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern float GameShare_nativeCallFloatFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern bool GameShare_nativeCallBoolFuncWithParam(string functionName, GameParam[] param,int count);
		
		[DllImport(GameUtil.GAME_PLATFORM)]
		private static extern void GameShare_nativeCallStringFuncWithParam(string functionName, GameParam[] param,int count,StringBuilder value);
	}
	
}


