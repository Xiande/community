//每个 待转的控件加 "lang"属性，不必设置值，将此js引用添加到页面或母版页中即可
var dict = null;
$(function () {
    //registerWords();
    //translate();
})
function translate() { 
    loadDict();
    $("[lang]").each(function () {
        switch (this.tagName.toLowerCase()) {
            case "input":
                $(this).val(__tr($(this).attr("lang")));
                break;
            case "html":
                break;

            default:
                $(this).text(__tr($(this).attr("lang")));
        }
    });
}

function __tr(src) {
    return (dict[src] || src);
}

function loadDict() {
	var lang = getCookie('CurLang');
	if(lang == null || lang=='null'){
		lang = CCConfig.SiteLanguage;
		    setCookie('CurLang',lang, 'd30');
	}
	else{
		CCConfig.SiteLanguage = lang;
	}
    $.ajax({
        async: false,
        type: "GET",
        url: CCConfig.L_Menu_BaseUrl + "/static/resource/" + CCConfig.SiteLanguage+ ".dict",
        success: function (msg) {
            dict = eval("(" + msg + ")");
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.statusText)
        }
    });
}

function registerWords() {
    $("[lang]").each(function () {
        switch (this.tagName.toLowerCase()) {
            case "input":
                $(this).attr("lang", $(this).val());
                break;
            case "html":
                break;
            default:
                $(this).attr("lang", $(this).text());
        }
    });
}
