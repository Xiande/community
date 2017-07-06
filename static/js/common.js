/// <reference path="jquery-1.7.1.min.js" />
String.prototype.format = function (args) {
    if (arguments.length > 0) {
        var result = this;
        if (arguments.length == 1 && typeof (args) == "object") {
            for (var key in args) {
                var reg = new RegExp("({" + key + "})", "g");
                result = result.replace(reg, args[key]);
            }
        }
        else {
            for (var i = 0; i < arguments.length; i++) {
                if (arguments[i] == undefined) {
                    return "";
                }
                else {
                    var reg = new RegExp("({[" + i + "]})", "g");
                    result = result.replace(reg, arguments[i]);
                }
            }
        }
        return result;
    }
    else {
        return this;
    }
}
String.prototype.trimStr = function (c) {
    if (c == null || c == "") {
        var str = this.replace(/^\/s*/, '');
        return str;
    }
    else {
        var reg = /,$/;
        var str = this.replace(reg, '');
        return str;
    }
};
String.prototype.subStr = function (length) {
    if (length != null && this.length > length) {
		return this.substring(0,length) + "...";    
    }else{
    	return this;
    }
};

String.prototype.encode = function () {
    return encodeURIComponent(this);
};
String.prototype.GuidStr = function () {
    var guid = "00000000-0000-0000-0000-000000000000"
    if (guid == this) {
        return "";
    } else {
        return this;
    }
};
String.prototype.Currency = function() {
	return this.replace(/\d+?(?=(?:\d{3})+$)/img, "$&,");
}
Number.prototype.ConvertCount = function () {
    if (this > 1000000) {
        return parseInt(this / 1000000) + "m";
    } else if (this > 1000) {
        return parseInt(this / 1000) + "k";
    } else {
        return this;
    }
}
Number.prototype.Currency = function() {
	return this.toString().replace(/\d+?(?=(?:\d{3})+$)/img, "$&,");
}

function GetRequestPara(paras) {
    var url = location.href;
    var paramStr = url.substring(url.indexOf('?') + 1, url.length).split('&');
    var j;
    var paramObj = {};
    for (var i = 0; j = paramStr[i]; i++) {
        paramObj[j.substring(0, j.indexOf('=')).toLowerCase()] = j.substring(j.indexOf('=') + 1, j.length);
    }

    var returnValue = paramObj[paras.toLowerCase()];

    if (typeof (returnValue) == "undefined") {
        return "";
    } else {
        return decodeURIComponent(returnValue.replace("#", ""));
    }
}
function getQueryString(name, searchStr) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var search = searchStr;
    if (search == null || search == "") {
        search = window.location.search.substr(1);
    }
    var r = search.match(reg);
    if (r != null)
        return unescape(r[2]);
    return "";
}

function getDateByDateStr(dateStr) {
    //dateStr Format   08/20/2014
    var ary = dateStr.split('/');
    var year = parseInt(ary[2]);
    var month = parseInt(ary[0], 10);
    var day = parseInt(ary[1], 10); //alert("year:" + year + "\nmonth:" + month + "\nday:" + day);
    return new Date(year, month - 1, day);
}

Date.prototype.format = function (format) {
    var o = {
        "M+": this.getMonth() + 1, //month 
        "d+": this.getDate(), //day 
        "h+": this.getHours(), //hour 
        "m+": this.getMinutes(), //minute 
        "s+": this.getSeconds(), //second 
        "q+": Math.floor((this.getMonth() + 3) / 3), //quarter 
        "S": this.getMilliseconds() //millisecond 
    }

    if (/(y+)/.test(format)) {
        format = format.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    }

    for (var k in o) {
        if (new RegExp("(" + k + ")").test(format)) {
            format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
        }
    }
    return format;
}

String.prototype.TransDate = function(){
	var d = new Date(this);
	return d.Trans();
}
Date.prototype.Trans = function(){
	//JavaScript函数：
	var minute = 1000 * 60;
	var hour = minute * 60;
	var day = hour * 24;
	var halfamonth = day * 15;
	var month = day * 30;
	var now = new Date().getTime();
	var diffValue = now - this;
	if(diffValue < 0){
	 //若日期不符则弹出窗口告之
	 //alert("结束日期不能小于开始日期！");
	 }
	var monthC =diffValue/month;
	var weekC =diffValue/(7*day);
	var dayC =diffValue/day;
	var hourC =diffValue/hour;
	var minC =diffValue/minute;
	if(monthC>=1){
	 result="发表于" + parseInt(monthC) + "个月前";
	 }
	 else if(weekC>=1){
	 result="发表于" + parseInt(weekC) + "周前";
	 }
	 else if(dayC>=1){
	 result="发表于"+ parseInt(dayC) +"天前";
	 }
	 else if(hourC>=1){
	 result="发表于"+ parseInt(hourC) +"个小时前";
	 }
	 else if(minC>=1){
	 result="发表于"+ parseInt(minC) +"分钟前";
	 }else
	 result="刚刚发表";
	return result;
}

function ajaxJsonCall(fullUrl, dataObj, callbackFunction) {
    $.ajax({
        type: 'get',
        url: fullUrl,
        data: dataObj,
        dataType: 'json',
        cache: false,
        success: function (data) {
            callbackFunction(data)
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            console.log("error :" + XMLHttpRequest.responseText);
            alert('There was an error in performing this operation.');
        }
    });
};

function startWaiting() {
    if ($("#dlgLoading").length <= 0) {
        $("#main-pannel").append('<div  id="dlgLoading" style="display: none;">Please waiting...</div>');
    }
    $("#dlgLoading").dialog({
        modal: true,
        open: function (event, ui) {
            $(".ui-dialog-titlebar-close", $(this).parent()).hide();
        }
    });
}

function stopWaiting() {
    if ($("#dlgLoading").length > 0) {
        $("#dlgLoading").dialog("close");
        $("#main-pannel").remove("#dlgLoading");
    }
}

function newGuid() {
    var guid = "";
    for (var i = 1; i <= 32; i++) {
        var n = Math.floor(Math.random() * 16.0).toString(16);
        guid += n;
        if ((i == 8) || (i == 12) || (i == 16) || (i == 20))
            guid += "-";
    }
    return guid;
}

function bindUserPresence() {
    var imgnID = 0;
    $(".up-container").each(function () {

        var itcode = $(this).attr("username");
        var name = $(this).attr("aname");
		var photoImgSrc = $(this).attr("photoImgSrc")
        var email = $(this).attr("aemail");
        var html = '<span class="ms-imnSpan">' +
        '<a href="/member/' + itcode + '"><img name="imnmark" title="" src="'+ photoImgSrc + '" alt="Available" style="margin-right:5px"></a>' +
        '<a class="ms-subtleLink" href="/member/' + itcode + '" target="_blank">' + name +
        '</a></span>';

        $(this).html(html);
        imgnID++;
    });

}
function SwitchTab(id) {
    $(".nav-tabs li").each(function () {
        $(this).removeClass("active");
    });
    $(".tab-pane").hide();
    $("#" + id).attr("class", "active");
    $("#con_" + id).show()
}

function CSSByNum(count) {
    if (count == null || count == 0) {
        return "a0";
    }
    else if (count > 0 && count < 20) {
        return "a19";
    }
    else if (count >= 20) {
        return "a20";
    }
}
function IndexInit() {
    Tags.GetTagsPage("tagspage");
    Questions.getQuestionPage("question");
    Miscellaneous.GetPopularSite("popularsite");
    Questions.GetTopDynamic("TopDynamic");
}

function goToSearchResult(obj, type){
	if(event.keyCode == 13){
		var key = obj.value;
		location.href="/questions?skey=" + key;
	}
}

function onSearch(obj, type){
	var key = obj.value;
	location.href="/questions?skey=" + key;
}

function getCookie(name){
	var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
	if(arr=document.cookie.match(reg))
		return unescape(arr[2]);
	else
		return null;
}

function setCookie(name,value,time){
	var strsec = getsec(time);
	var exp = new Date();
	exp.setTime(exp.getTime() + strsec*1);
	document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
}
function getsec(str){
	var str1=str.substring(1,str.length)*1;
	var str2=str.substring(0,1);
	if (str2=="s"){
		return str1*1000;
	}else if (str2=="h")	{
		return str1*60*60*1000;
	}else if (str2=="d"){
		return str1*24*60*60*1000;
	}
}

function changeLanguage(lang){
	setCookie('CurLang',lang, 'd30');
	//CCConfig.SiteLanguage=lang;
	//translate();
	window.location.reload();
}

function ccAlert(dictKey, msg){
	if ( dict == null){
		loadDict();
	}
	
	var str = dict[dictKey];
	if (msg != undefined)
		str += msg;
	
	alert(str);
}
function GetPlural(key,num){
	if(num <= 1){
		return dict[key];
	} else {
		return dict[key + "s"];
	}
}
function bindFixedCount(count) {
	$("#fixedQuestionCount").text(count.Currency());
}
function BindMenuLink() {
    var lang = getCookie('CurLang');
    if (lang == "zh-cn") {
        $("#coworkindex").attr("href", "http://cowork.cn.lenovo.com/ ");
        $("#applypage").attr("href", "http://cowork.cn.lenovo.com/SitePages/Apply%20for%20Sites.aspx");
        $("#sitefaq").attr("href", "http://cowork.cn.lenovo.com/SitePages/Site%20FAQ.aspx");
        $("#helpcenter").attr("href", "http://cowork.cn.lenovo.com/promotions/sphelp_cn/SitePages/Home.aspx ");
    } else {
        $("#coworkindex").attr("href", "http://cowork.us.lenovo.com");
        $("#applypage").attr("href", "http://cowork.us.lenovo.com/SitePages/Apply%20for%20Sites.aspx");
        $("#sitefaq").attr("href", "http://cowork.us.lenovo.com/SitePages/Site%20FAQ.aspx");
        $("#helpcenter").attr("href", "http://cowork.us.lenovo.com/promotions/help_en/SitePages/Home.aspx");
    }
}
function IsURL(str_url){
	var strRegex = "^((https|http):\/\/)"  
		     + "(((([0-9]|1[0-9]{2}|[1-9][0-9]|2[0-4][0-9]|25[0-5])[.]{1}){3}([0-9]|1[0-9]{2}|[1-9][0-9]|2[0-4][0-9]|25[0-5]))" // IP>形式的URL- 199.194.52.184  
		     + "|"  
		     + "([0-9a-zA-Z\u4E00-\u9FA5\uF900-\uFA2D-]+[.]{1})+[a-zA-Z-]+)" // DOMAIN（域名）形式的URL  
		     + "(:[0-9]{1,4})?" // 端口- :80  
		     + "((/?)|(/[0-9a-zA-Z_!~*'().;?:@&=+$,%#-]+)+/?){1}";  
	var RegUrl = new RegExp();
    RegUrl.compile(strRegex);
    if (!RegUrl.test(str_url) && str_url != ""){
        return false;
    }
    return true;
}
function showLoading(){
	var h = $(".ind-lef").height();//.css("height");
	$("#waitingCover").height(h + 60);//.css("heigth", h);
	$("#loading").show();
}
Array.prototype.filter = Array.prototype.filter || function(fun /*, thisArg */){
    "use strict";
    if (this === void 0 || this === null)
        throw new TypeError();
    var t = Object(this);
    var len = t.length >>> 0;
    if (typeof fun !== "function")
        throw new TypeError();
    var res = [];
    var thisArg = arguments.length >= 2 ? arguments[1] : void 0;
    for (var i = 0; i < len; i++) {
        if (t.hasOwnProperty(i)) {
            var val = t[i];
            if (fun.call(thisArg, val, i, t))
                res.push(val);
        }
    }
    
    return res;
};

Array.prototype.indexOf = Array.prototype.indexOf || function(elt /*, from*/)
  {
    var len = this.length >>> 0;
    var from = Number(arguments[1]) || 0;
    from = (from < 0)
         ? Math.ceil(from)
         : Math.floor(from);
    if (from < 0)
      from += len;
    for (; from < len; from++)
    {
      if (from in this &&
          this[from] === elt)
        return from;
    }
    return -1;
  };
