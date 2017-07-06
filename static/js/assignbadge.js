var viewModel = {};
function initializePeoplePicker(peoplePickerElementId) {

    // Create a schema to store picker properties, and set the properties.
    var schema = {};
    schema['PrincipalAccountType'] = 'User';
    schema['SearchPrincipalSource'] = 15;
    schema['ResolvePrincipalSource'] = 15;
    schema['AllowMultipleValues'] = true;
    schema['MaximumEntitySuggestions'] = 50;
    schema['Width'] = '100%';


    // Render and initialize the picker. 
    // Pass the ID of the DOM element that contains the picker, an array of initial
    // PickerEntity objects to set the picker value, and a schema that defines
    // picker properties.
    this.SPClientPeoplePicker_InitStandaloneControlWrapper(peoplePickerElementId, null, schema);
    $('#peoplePickerDiv_TopSpan').addClass('form-control');
    $('#peoplePickerDiv_TopSpan_InitialHelpText').css('padding-top', '8px');
    $('#peoplePickerDiv_TopSpan_InitialHelpText').text(dict['UserPickHelpText']);
    //$('#peoplePickerDiv_TopSpan').css('height','30px');
    //$('#peoplePickerDiv_TopSpan').css('padding','5px 0px 0px 0px');

}
function bindBadges(targetId){
	Badges.getAllBadges(function(data){
		if(!data.HasPermission){
			if(confirm(dict["NoAssignPermission"])){
				Badges.applyPermission(closePage());
			}
			else{
				closePage();
			}
		}
		else
		{
			viewModel = ko.mapping.fromJS(data);
		    //ko.options.useOnlyNativeEvents = true;
		    ko.applyBindings(viewModel, document.getElementById(targetId));
		    
		    $(".badgeimg").hover(function(){
		    	var cls1 = ".badgemessage-" + CCConfig.SiteLanguage;
		    	var cls2 = ".badgemessagebg-" + CCConfig.SiteLanguage;
		    	$(this).parent().children(cls1).show();
		    	$(this).parent().children(cls2).show();

		    },function(){
		    	$(this).parent().children(".badgemessage-" + CCConfig.SiteLanguage).hide();
		    	$(this).parent().children(".badgemessagebg-" + CCConfig.SiteLanguage).hide();
		    })
		    

	    }
	});
}
function closeCurrent(){
	var h = $(".ind-lef").height();//.css("height");
	$("#SuggestCoverlayer").height(h + 60);//.css("heigth", h);
	$("#closeaskg").show();
	

/*
	var source = getQueryString("Source");
	var isDlg = getQueryString("IsDlg");
	if(isDlg != null && isDlg == "1"){
		window.frameElement.commonModalDialogClose();
	}
	else{
	    if(source == null || source == ""){
	    	window.location = "index.aspx";
	    }
	    else{
	    	window.location = source;
	    }
    }
    */
}
function closePage(){
	var source = getQueryString("Source");
	var isDlg = getQueryString("IsDlg");
	if(isDlg != null && isDlg == "1"){
		window.frameElement.commonModalDialogClose();
	}
	else{
	    if(source == null || source == ""){
	    	window.location = "index.aspx";
	    }
	    else{
	    	window.location = source;
	    }
    }

}
function reSend(){
	window.location.reload();
}
function assignBadge(){
	
	var peoplePicker = this.SPClientPeoplePicker.SPClientPeoplePickerDict.peoplePickerDiv_TopSpan;

    // Get information about all users.
    var users = peoplePicker.GetAllUserInfo();
    var inputUsers= [];
    var inputNames=[];
    /*
    var strs = users[0].Key.split("|")
    var user = strs[strs.length - 1];
    inputUsers.push(user);
    */
    var hasSelf = false;
    for(var i = 0; i < users.length; i++){
    	var strs = users[i].Key.split("|")
    	var user = strs[strs.length - 1];
    	if(user == viewModel.CurUserLoginName()){
    		hasSelf = true;
    	}
    	else if(inputUsers.indexOf(user) < 0){
			inputUsers.push(user);
			inputNames.push(users[i].DisplayText);
		}
    }
    
    if(hasSelf){
    	ccAlert("SelfAssigned");
    }
    
    if (inputUsers.length == 0){
    	ccAlert("UserNotNull");
    	return
    }
    
	var selectedBadges = [];
	var allBadges = ko.mapping.toJS(viewModel.AllBadges);
	for(var i = 0; i < allBadges.length; i++){
		for(var j = 0; j < allBadges[i].BadgesList.length; j++){
			if(allBadges[i].BadgesList[j].Selected){
				selectedBadges.push(allBadges[i].BadgesList[j]);
			}
		}
	}
	
	if(selectedBadges.length == 0){
		ccAlert("SelectedBadgesIsNull");
    	return
	}
	
	if( $("#message").val().trim() == ""){
		ccAlert("BadgeReasonIsNull");
    	return;
	}
	
	var badgeWallUrl = CCConfig.BadgeWallUrl + "?accountname=" ;
	
    var data = {
    	UserLoginNames:inputUsers,
    	Badges:selectedBadges,
    	Message:$("#message").val(),
    	BadgeWallUrl:badgeWallUrl
    }
    
    var html="";
    for(var i = 0; i< inputUsers.length; i++){
    	html += "<i class='glyphicon glyphicon-user mr5'></i><a class='mr10' target='blank' href='"+ badgeWallUrl  + inputUsers[i] +"'>"+ inputNames[i] +"</a>";
    }
    html = $("#badgeWallMsg").html() + html;
    $("#badgeWallMsg").html(html);
    
    Badges.assignBadge(data,closeCurrent);	
}