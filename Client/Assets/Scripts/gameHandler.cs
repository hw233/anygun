using UnityEngine;
using System;
using System.Collections.Generic;
using game;

public class gameHandler : MonoBehaviour {

	public static bool _SdkInitSuccess;

	public static PayInfo _PayInfo;

	Dictionary<QuestKind, string> _QuestTypeMirror;

	// Use this for initialization
	void Start () {
		TransferRate._Inst.Send("Init SDK Begin");
		_QuestTypeMirror = new Dictionary<QuestKind, string>();
		_QuestTypeMirror[QuestKind.QK_Main] = ((int)TaskType.MAIN_LINE).ToString();
		_QuestTypeMirror[QuestKind.QK_Daily] = ((int)TaskType.DAILY).ToString();
		_QuestTypeMirror[QuestKind.QK_Profession] = ((int)TaskType.OTHER).ToString();
		_QuestTypeMirror[QuestKind.QK_Sub] = ((int)TaskType.BRANCH_LINE).ToString();
		_QuestTypeMirror[QuestKind.QK_Tongji] = ((int)TaskType.ACTIVITY).ToString();

		string appKey = "FDBE1788-C7E2-94A9-0B79-7B4B418F5871";
		string appSecret = "8bd1bd4e5ad4ae603039a6c6e3f41ecb";
		string privateKey = "0FBA2600D9720A14793946F640CF8708";
		string oauthLoginServer = "http://10.10.10.253:8080/AnyLoginOauth/";

		Game.getInstance().init(appKey, appSecret, privateKey, oauthLoginServer);

		GameUser.getInstance().setListener(this, "UserExternalCall");

		GameIAP.getInstance().setListener(this, "IAPExternalCall");

		CommonEvent.OnAppPause += OnAppPause;
		CommonEvent.OnAppResume += OnAppResume;
		CommonEvent.OnAccountChange += OnAccountChange;//Login regist logout
		CommonEvent.OnQuestFinish += OnQuestFinish;
		CommonEvent.OnQuestFail += OnQuestFail;
		CommonEvent.OnQuestStart += OnQuestStart;
		CommonEvent.OnException += OnException;
		CommonEvent.OnPurchase += OnPurchase;
		CommonEvent.OnUseItem += OnUseItem;
		CommonEvent.OnRewardVirtualCash += OnRewardVirtualCash;

#if UNITY_IOS || UNITY_IPHONE
        OnAppResume();
#endif
	}

    public void Reinit()
    {
        Game.getInstance().release();
        string appKey = "FDBE1788-C7E2-94A9-0B79-7B4B418F5871";
        string appSecret = "8bd1bd4e5ad4ae603039a6c6e3f41ecb";
        string privateKey = "0FBA2600D9720A14793946F640CF8708";
        string oauthLoginServer = "http://10.10.10.253:8080/AnyLoginOauth/";
        Game.getInstance().init(appKey, appSecret, privateKey, oauthLoginServer);
        GameUser.getInstance().setListener(this, "UserExternalCall");
        GameIAP.getInstance().setListener(this, "IAPExternalCall");
    }

	void OnAppPause()
	{
		GameAnalytics.getInstance().stopSession();
	}

	void OnAppResume()
	{
		GameAnalytics.getInstance().startSession();
	}

	public void ReInit()
	{
		Game.getInstance ().release ();
		string appKey = "FDBE1788-C7E2-94A9-0B79-7B4B418F5871";
		string appSecret = "8bd1bd4e5ad4ae603039a6c6e3f41ecb";
		string privateKey = "0FBA2600D9720A14793946F640CF8708";
		string oauthLoginServer = "http://10.10.10.253:8080/AnyLoginOauth/";
		
		Game.getInstance().init(appKey, appSecret, privateKey, oauthLoginServer);
		
		GameUser.getInstance().setListener(this, "UserExternalCall");
		
		GameIAP.getInstance().setListener(this, "IAPExternalCall");
	}

	void OnAccountChange(CommonEvent.DefineAccountOperate type)
	{
        Dictionary<string, string> paramMap = new Dictionary<string, string>();
        paramMap["Account_Id"] = GameUser.getInstance().getUserID();
        paramMap["Account_Name"] = GamePlayer.Instance.InstName;
        paramMap["Account_Type"] = Convert.ToString((int)AccountType.ANONYMOUS);
        paramMap["Account_Level"] = GamePlayer.Instance.GetIprop(PropertyType.PT_Level).ToString();
        paramMap["Account_Operate"] = Convert.ToString((int)type);
        paramMap["Account_Gender"] = Convert.ToString((int)AccountGender.UNKNOWN);
        paramMap["Server_Id"] = GameManager.ServId_.ToString();
        paramMap["Account_Age"] = "18";
        GameParam param1 = new GameParam(paramMap);
        GameAnalytics.getInstance().callFuncWithParam("setAccount", param1);

        if (type == CommonEvent.DefineAccountOperate.LOGIN)
        {
            Dictionary<string, string> dic = new Dictionary<string, string>();
            dic["User_Id"] = GameUser.getInstance().getUserID();
            dic["Role_Id"] = GamePlayer.Instance.InstId.ToString();
            dic["Role_Name"] = GamePlayer.Instance.InstName;
            GameAdTracking.getInstance().onLogin(dic);

            if (GamePlayer.Instance.isCreate)
            {
                if (GameAdTracking.getInstance().isFunctionSupported("onCreateRole"))
                    GameAdTracking.getInstance().trackEvent("onCreateRole", dic);
            }
        }
        else if (type == CommonEvent.DefineAccountOperate.REGISTER)
        {
            GameAdTracking.getInstance().onRegister(GameUser.getInstance().getUserID());
        }

		if(GamePlayer.Instance.InstId == 0)
			return;

		if(string.IsNullOrEmpty(GamePlayer.Instance.InstName))
			return;

		if(GamePlayer.Instance.GetIprop(PropertyType.PT_Level) == 0)
			return;

		if(GamePlayer.Instance.createTime_ == 0)
			return;

		if(GameUser.getInstance().isFunctionSupported("submitLoginGameRole"))
		{
			Dictionary<string, string> map = new Dictionary<string, string>();
			map["roleId"]= GamePlayer.Instance.InstId.ToString();
			map["roleName"]= GamePlayer.Instance.InstName;
			map["roleLevel"]= GamePlayer.Instance.GetIprop(PropertyType.PT_Level).ToString();
			map["zoneId"]= GameManager.ServId_.ToString();
            map["zoneName"] = GameManager.ServName_;
			map["roleCTime"]= Convert.ToString(GamePlayer.Instance.createTime_);
            map["balance"] = GamePlayer.Instance.GetIprop(PropertyType.PT_MagicCurrency).ToString();
            map["partyName"] = GuildSystem.Mguild == null ? "" : GuildSystem.Mguild.guildName_;
            map["vipLevel"] = GamePlayer.Instance.GetIprop(PropertyType.PT_VipLevel).ToString();
			string sdType = "";
			switch(type)
			{
			case CommonEvent.DefineAccountOperate.LOGIN:
				if(GamePlayer.Instance.isCreate)
					sdType = "2";
				else
					sdType = "1";
                if (GlobalValue.channelID.Equals("000003"))
                    map["roleLevelMTime"] = crtTimeStamp();
                else
                    map["roleLevelMTime"] = "-1";
				break;
			case CommonEvent.DefineAccountOperate.LOGOUT:
				sdType = "4";
                if (GlobalValue.channelID.Equals("000003"))
                    map["roleLevelMTime"] = crtTimeStamp();
                else
                    map["roleLevelMTime"] = "-1";
				break;
			default:
				sdType = "1";
				map["roleLevelMTime"]= crtTimeStamp();
				break;
			}
			map["dataType"]= sdType;
			GameParam param = new GameParam(map);
			GameUser.getInstance().callFuncWithParam("submitLoginGameRole",param);
			ClientLog.Instance.Log(GamePlayer.Instance.InstId.ToString() + " " + GamePlayer.Instance.InstName + " " + GamePlayer.Instance.GetIprop(PropertyType.PT_Level).ToString() + " " + GamePlayer.Instance.createTime_.ToString());
		}
	}
	
	string crtTimeStamp()
	{
		return ConvertDateTimeInt(DateTime.Now).ToString();
	}

	int ConvertDateTimeInt(System.DateTime time)
	{
		DateTime startT = TimeZone.CurrentTimeZone.ToLocalTime(new DateTime(1970, 1, 1));
		return (int)(time - startT).TotalSeconds;
	}

	void OnQuestStart(int questid)
	{
		QuestData qData = QuestData.GetData(questid);
		if(qData != null)
		{
			Dictionary<string, string> map = new Dictionary<string, string>();
			map["Task_Id"] = questid.ToString();
			map["Task_Type"] = _QuestTypeMirror[qData.questKind_];
			GameParam param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("startTask", param);
		}
	}

	void OnQuestFinish(int questid)
	{
		GameAnalytics.getInstance().callFuncWithParam("finishTask", new GameParam(questid.ToString()));
	}

	void OnQuestFail(int questid)
	{
		Dictionary<string, string> map = new Dictionary<string, string>();
		map["Task_Id"] = questid.ToString();
		map["Fail_Reason"] = "GiveUp";
		GameParam param = new GameParam(map);
		GameAnalytics.getInstance().callFuncWithParam("failTask", param);
	}

	void OnException(string msg)
	{
		//GameAnalytics.getInstance().logError(Guid.NewGuid().ToString(), msg);
	}

	void OnPurchase(int itemid, int stack, int price)
	{
		ItemData iData = ItemData.GetData(itemid);
		if(iData != null)
		{
			Dictionary<string, string> map = new Dictionary<string, string>();
			map["Item_Id"] = itemid.ToString();
			map["Item_Type"] = iData.subType_.ToString();
			map["Item_Count"] = stack.ToString();
			map["Virtual_Currency"] = price.ToString();
			map["Currency_Type"] = Game.getInstance().getChannelId();
			GameParam param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onPurchase", param);
		}
	}

	void OnUseItem(int itemid, int stack)
	{
		ItemData iData = ItemData.GetData(itemid);
		if(iData != null)
		{
			Dictionary<string, string> map = new Dictionary<string, string>();
			map["Item_Id"] = itemid.ToString();
			map["Item_Type"] = iData.subType_.ToString();
			map["Item_Count"] = stack.ToString();
			map["Use_Reason"] = "byDefault";
			GameParam param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onUse", param);
		}
	}

	void OnRewardVirtualCash(int count, string reason)
	{
		Dictionary<string, string> map = new Dictionary<string, string>();
		map["Item_Count"] = count.ToString();
		map["Use_Reason"] = reason;
		GameParam param = new GameParam(map);
		GameAnalytics.getInstance().callFuncWithParam("onReward", param);
	}

	void UserExternalCall(string msg)
	{
		Dictionary<string,string> dic = GameUtil.stringToDictionary (msg);
		int code = Convert.ToInt32(dic["code"]);
		string result = dic["msg"];
		ClientLog.Instance.Log("UserExternalCall( code: "+  + code + "  result: " + result +")");
		switch(code)
		{
		case (int)UserActionResultCode.kInitSuccess://初始化SDK成功回调
            if(string.IsNullOrEmpty(GlobalValue.channelID))
                GlobalValue.channelID = game.Game.getInstance().getChannelId();
//            string custom = game.Game.getInstance().getCustomParam();
//            if(!string.IsNullOrEmpty(custom))
//            {
//                string[] centerAndcdnAddr = custom.Split(new char[]{';'}, StringSplitOptions.RemoveEmptyEntries);
//                if(centerAndcdnAddr.Length == 2)
//                {
//                    GlobalValue.centerservhost = centerAndcdnAddr[0];
//					GlobalValue.cdnservhost = string.Format("{0}/{1}/", centerAndcdnAddr[1], XyskAndroidAPI.getPackageVersion());
//					Debug.Log(GlobalValue.cdnservhost);
//                }
//            }

			if(GlobalValue.IsDebugMode)
			{
				CommonEvent.OnAppPause -= OnAppPause;
				CommonEvent.OnAppResume -= OnAppResume;
				CommonEvent.OnAccountChange -= OnAccountChange;//Login regist logout
				CommonEvent.OnQuestFinish -= OnQuestFinish;
				CommonEvent.OnQuestFail -= OnQuestFail;
				CommonEvent.OnQuestStart -= OnQuestStart;
				CommonEvent.OnException -= OnException;
				CommonEvent.OnPurchase -= OnPurchase;
				CommonEvent.OnUseItem -= OnUseItem;
				CommonEvent.OnRewardVirtualCash -= OnRewardVirtualCash;
			}
			_SdkInitSuccess = true;
			TransferRate._Inst.Send("Init SDK End Success");
			break;
		case (int)UserActionResultCode.kInitFail://初始化SDK失败回调
			ApplicationEntry.Instance.PostSocketErr(1234);
			TransferRate._Inst.Send("Init SDK End Failed");
			break;
		case (int)UserActionResultCode.kLoginSuccess://登陆成功回调
            GameManager.Instance.loginInfo_ = new COM_LoginInfo();
			GameManager.Instance.loginInfo_.username_ = GameUser.getInstance().getPluginId() + "=" + GameUser.getInstance().getUserID();
           // GameManager.Instance.loginInfo_.username_ = GameUser.getInstance().getUserID();
            GameManager.Instance.loginInfo_.password_ = Application.platform.ToString();
            GameManager.Instance.loginInfo_.version_ = Configure.VersionNumber;
            GameManager.Instance._Account = GameManager.Instance.loginInfo_.username_;
            GameManager.Instance.loginInfo_.sessionkey_ = Configure.Sessionkey;
#if UNITY_IOS || UNITY_IPHONE
            GameManager.Instance.loginInfo_.idfa_ = XyskIOSAPI.GetIDFA();
#elif UNITY_ANDROID
            GameManager.Instance.loginInfo_.mac_ = XyskAndroidAPI.getMacAndroid();
#endif
			game.GameParam param = new game.GameParam((int)game.ToolBarPlace.kToolBarMidRight);
			if(game.GameUser.getInstance().isFunctionSupported("showToolBar"))
				game.GameUser.getInstance().callFuncWithParam("showToolBar", param);

            CommonEvent.ExcuteAccountChange(CommonEvent.DefineAccountOperate.REGISTER);
			TransferRate._Inst.Send("Login SDK End Success");
			break;
		case (int)UserActionResultCode.kLoginNetworkError://登陆网络出错回调
		case (int)UserActionResultCode.kLoginFail://登陆失败回调
			ApplicationEntry.Instance.PostSocketErr(2345);
			TransferRate._Inst.Send("Login SDK End Failed");
			break;
        case (int)UserActionResultCode.kLoginCancel://登陆取消回调
            break;
		case (int)UserActionResultCode.kLogoutSuccess://登出成功回调
			if(StageMgr.Loading)
				StageMgr.OnSceneLoaded += SceneLoaded;
            else if (!string.IsNullOrEmpty(StageMgr.Scene_name) && !StageMgr.Scene_name.Equals(GlobalValue.StageName_ReLoginScene))
				ApplicationEntry.Instance.PostSocketErr(2333);
			break;
		case (int)UserActionResultCode.kLogoutFail://登出失败回调
			break;
		case (int)UserActionResultCode.kPlatformEnter://平台中心进入回调
			break;
		case (int)UserActionResultCode.kPlatformBack://平台中心退出回调
			break;
		case (int)UserActionResultCode.kPausePage://暂停界面回调
			break;
		case (int)UserActionResultCode.kExitPage://退出游戏回调
			CommonEvent.ExcuteAccountChange(CommonEvent.DefineAccountOperate.LOGOUT);
			Application.Quit();
			break;
		case (int)UserActionResultCode.kAntiAddictionQuery://防沉迷查询回调
			break;
		case (int)UserActionResultCode.kRealNameRegister://实名注册回调
			break;
		case (int)UserActionResultCode.kAccountSwitchSuccess://切换账号成功回调
            if (StageMgr.Loading)
                StageMgr.OnSceneLoaded += SceneLoaded;
            else if (!string.IsNullOrEmpty(StageMgr.Scene_name) && !StageMgr.Scene_name.Equals(GlobalValue.StageName_ReLoginScene))
                ApplicationEntry.Instance.PostSocketErr(2333);
            else
                game.GameUser.getInstance().login();
			break;
		case (int)UserActionResultCode.kAccountSwitchFail://切换账号成功回调
			GameUser.getInstance().login();
			break;
		default:
			break;
		}
        CommonEvent.ExcuteUserExternal(code);
	}

	void SceneLoaded(string sceneName)
	{
		StageMgr.OnSceneLoaded -= SceneLoaded;
		ApplicationEntry.Instance.PostSocketErr(2333);
	}

	public void OrderOk(string orderId, int shopId)
	{
		Dictionary<string,string> map;
		GameParam param;

		ShopData sData = ShopData.GetData(shopId);
		map = new Dictionary<string,string>();
		map["Order_Id"]= orderId;
		map["Product_Name"]= sData == null? "": sData._Name;
		map["Currency_Amount"]= sData == null? "": sData._Price.ToString();
		map["Currency_Type"]= "CNY";
		map["Payment_Type"]= Game.getInstance().getChannelId();
		map["Virtual_Currency_Amount"]= sData == null? "": sData._Num.ToString();
		param = new GameParam(map);
		GameAnalytics.getInstance().callFuncWithParam("onChargeRequest",param);

		//这里统计支付完成
		PopText.Instance.Show(LanguageManager.instance.GetValue("paySuccess"), PopText.WarningType.WT_Tip);
		param = new GameParam(orderId);
		GameAnalytics.getInstance().callFuncWithParam("onChargeSuccess",param);
        map = new Dictionary<string, string>();
        map["User_Id"] = GamePlayer.Instance.InstId.ToString();
        map["Order_Id"] = orderId;
        map["Currency_Amount"] = sData == null ? "" : sData._Price.ToString();
        map["Currency_Type"] = "CNY";
        map["Payment_Type"] = Game.getInstance().getChannelId();
		GameAdTracking.getInstance().onPay(map);
	}

	void IAPExternalCall(string msg)
	{
		Dictionary<string,string> dic = GameUtil.stringToDictionary (msg);
		int code = Convert.ToInt32(dic["code"]);
		string result = dic["msg"];
		ClientLog.Instance.Log("IAPExternalCall( code: "+  + code + "  result: " + result +")");

		Dictionary<string,string> map;
		GameParam param;

		string orderId = GameIAP.getInstance().getOrderId(GameIAP.getInstance().getPluginId()[0]);

		if(code != (int)PayResultCode.kPaySuccess)
		{
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Product_Name"]= _PayInfo == null? "": orderId;
			map["Currency_Amount"]= _PayInfo == null? "": _PayInfo._PayPrice;
			map["Currency_Type"]= "CNY";
			map["Payment_Type"]= Game.getInstance().getChannelId();
			map["Virtual_Currency_Amount"]= _PayInfo == null? "": _PayInfo._PayVirtual;
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeRequest",param);
		}

		switch(code)
		{
		case (int)PayResultCode.kPaySuccess://支付成功回调

			break;
		case (int)PayResultCode.kPayFail://支付失败回调
			PopText.Instance.Show(LanguageManager.instance.GetValue("payFail"), PopText.WarningType.WT_Warning);
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "PayFail";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
		case (int)PayResultCode.kPayCancel://支付取消回调
			PopText.Instance.Show(LanguageManager.instance.GetValue("payCancel"), PopText.WarningType.WT_Warning);
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "PayCancel";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
		case (int)PayResultCode.kPayNetworkError://支付超时回调
			PopText.Instance.Show(LanguageManager.instance.GetValue("payNetworkErr"), PopText.WarningType.WT_Warning);
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "payNetworkErr";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
		case (int)PayResultCode.kPayProductionInforIncomplete://支付信息不完整
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "payIncompleteInfo";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
			/**
        * 新增加:正在进行中回调
        * 支付过程中若SDK没有回调结果，就认为支付正在进行中
        * 游戏开发商可让玩家去判断是否需要等待，若不等待则进行下一次的支付
        */
		case (int)PayResultCode.kPayNowPaying:
			GameIAP.getInstance().resetPayState();
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "otherPaying";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
		default:
			map = new Dictionary<string,string>();
			map["Order_Id"]= orderId;
			map["Fail_Reason"]= "otherReason";
			param = new GameParam(map);
			GameAnalytics.getInstance().callFuncWithParam("onChargeFail",param);
			break;
		}

        _PayInfo = null;
	}

    public static void PayProduct(int shopId)
    {
        ShopData sData = ShopData.GetData(shopId);
        if (sData == null)
        {
            ClientLog.Instance.Log(" shot id is null ");
            return;
        }

        Dictionary<string, string> mProductionInfo = new Dictionary<string, string>();
		mProductionInfo["Product_Price"] = sData._Price.ToString();
        int Server_Id = GameManager.ServId_;
#if UNITY_IPHONE
		//GlobalValue.Get(Constant.C_IosPayServerId, out Server_Id);
		mProductionInfo["Product_Id"] =sData._IosShopid;
		mProductionInfo["Server_Id"] = Server_Id.ToString();
#elif UNITY_ANDROID

		mProductionInfo["Product_Id"] =shopId.ToString();
		Debug.Log(" pay code : " + shopId.ToString());

		mProductionInfo["Server_Id"] = Server_Id.ToString();
#endif
        string payDes = sData._Name;
        if (GlobalValue.channelID.Equals("500002"))
            payDes = sData._Id.ToString();
        mProductionInfo["Product_Name"] = payDes;
		mProductionInfo["Product_Count"] = "1";//sData._Num.ToString();
		mProductionInfo["Role_Id"] = GamePlayer.Instance.InstId.ToString();
        mProductionInfo["Role_Name"] = GamePlayer.Instance.InstName;
        mProductionInfo["Role_Grade"] = GamePlayer.Instance.GetIprop(PropertyType.PT_Level).ToString();
        mProductionInfo["Role_Balance"] = GamePlayer.Instance.GetIprop(PropertyType.PT_MagicCurrency).ToString();
        mProductionInfo["Coin_Name"] = LanguageManager.instance.GetValue("");
        mProductionInfo["Coin_Rate"] = "10";
        mProductionInfo["Vip_Level"] = GamePlayer.Instance.GetIprop(PropertyType.PT_VipLevel).ToString();
        mProductionInfo["Server_Name"] = GameManager.ServName_;
		mProductionInfo["EXT"] = "pay";

        List<string> idArrayList = game.GameIAP.getInstance().getPluginId();
        if (idArrayList == null)
        {
            ClientLog.Instance.Log(" idArrayList is null ");
            return;
        }
        if (idArrayList.Count == 1)
            game.GameIAP.getInstance().payForProduct(mProductionInfo);

        _PayInfo = new PayInfo();
        _PayInfo._ProductName = sData._Name;
		_PayInfo._PayPrice = sData._Price.ToString();
        _PayInfo._PayVirtual = sData._Num.ToString();
    }

    void PhoneBindHandler(int type, int code, string message)
    {
        switch (type)
        {
            case 0: // 申请短信认证码
                if (code == 0 && string.IsNullOrEmpty(message))
                    PopText.Instance.Show(LanguageManager.instance.GetValue("PhoneCodeSended"), PopText.WarningType.WT_Tip);
                break;
            case 1: // 检查账号是否已认证
                break;
            case 2: // 检查账号与手机号是否匹配
                break;
            case 3: // 账号认证绑定
                if (code == 0 && string.IsNullOrEmpty(message))
                {
                    PopText.Instance.Show(LanguageManager.instance.GetValue("EN_PboneNumberSuccess"), PopText.WarningType.WT_Tip);
                    NetConnection.Instance.verificationSMS(GameManager.Instance.mobileNum, "");
                }
                break;
            case 4: // 账号认证解绑定
                break;
            default:
                break;
        }
    }
}

public class PayInfo
{
	public string _ProductName;
	public string _PayPrice;
	public string _PayVirtual;
}
