function closeWindow() {
    window.frameElement.cancelPopUp();
    return false;
}

function keyUp(obj) {
    //if ($(obj).val().length > 0) {
    //    obj.size = $(obj).val().length;
    //}
}

function keyPress(event) {
    if (event.keyCode == 13) {
        event.Handled = true;
        checkTag();
        
        if(event.preventDefault){
			event.preventDefault();
		}
		else{
			event.returnValue = false;
		}
    }
}

function checkTag() {
    //var reg = /^([.a-zA-Z0-9_-])+@([.a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/;
    var input = $("#inputTag");
    var content = input.val();
    var tags = content.split(",");

    for (var i = 0; i < tags.length; i++) {
        if (tags[i].trim() == "") {
            continue;
        }

       	if(tags[i].trim().length > 20){
			ccAlert('TagLengthValidation');
			return false;
		}
		
		if($(".tagItem").length >= 4){
	    	ccAlert('TagsTotalCountValidation');	
			return false;
	    }
	    
        addTag(tags[i].trim());
    }
    
    input.val("");
    
    return true;
}

function addTag(tag) {
   /*
    var html = '<span class="checkedItem">';
    html += '<span class="tagItem">' + tag + '</span>'
    html += '<a onclick="removeTag(this);" style="cursor:pointer;"> X </a>';
    html += '</span>';
    $("#checkedTags").append(html);
    */
    var html = '<li class="checkedItem"><a class="tagItem">'+  tag +'</a>';
    html += '<span onclick="removeTag(this);" style="cursor:pointer;"> X </span>';
    html += '</li>';
    $(html).insertBefore("#lastTag");

}

function generateTags() {
    var hideVal = "";
    $(".checkedItem").each(function (index, element) {
        //if (!$(element).hasClass("invalidMail")) {
        var tag = $(element).children(".tagItem").eq(0).text();//.toLowerCase();
        if (hideVal.indexOf(tag + ",") < 0) {
            hideVal += tag + ",";
        }
    });

    $("#hiddenTags").val(hideVal);
}

function removeTag(obj) {
    $(obj).parent().remove();
}

function recommendTags(obj) {
    var title = obj.value;
    //checkLanguageType(title);
    Tags.recommendTags(title);
}

function submitQuestion() {

    if ($("#txtTitle").val() == "") {
        ccAlert("TitleIsNull");
        return false;
    }

    if (CKEDITOR.instances.editor01.getData() == "") {
        ccAlert("ContentIsNull");
        return false;
    }
	
     if (!checkTag()){
     	return false;
     }

    if ($(".checkedItem").length == 0 || $(".checkedItem").length > 4) {
        $("#hiddenTags").val("");
        ccAlert("TagNullOrOverLimitationMessage");
        return false;
    }
    else {
        generateTags();
    }
	
	var title = $("#txtTitle").val();
	var lang = checkLanguageType(title);
    var question = {
        Title:title,
        Content: CKEDITOR.instances.editor01.getData(),
        Tags: $("#hiddenTags").val(),
        NoticeEmail: $("#ckEmailNotice")[0].checked,
        LanguageType: lang
    }

    Questions.addQuestion(question);

    return true;
}				

function checkLanguageType(title){
	var lang = $("#langselect").val();
	if (lang == "en"){
		var cnChar = matchChinese(title);
		if (cnChar != ""){
			var msg = dict["ChangeLanguageConfirm"];
			msg = msg.format(cnChar);
			 
			if (confirm(msg)){
				$("#langselect").val("cn");
				lang = "cn";
			}
		}
	}
	
	return lang;
}

function matchChinese(str){   
   //var reg = /[\u4E00-\u9FA5\uF900-\uFA2D\uFF00-\uFFEF]/;
   //var reg2 = /[\uFF00-\uFFEF]/;
   var reg = /[^\x00-\xff]+/;
   var index = str.search(reg);
   if (index >=0)
   	 return str[index];
   
   return "";
}
