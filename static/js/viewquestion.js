var viewQuestionModel = {};
//var Cowork = window.Cowork || {};

function bindData(data){
	
	if (data == null){
		$("#content-container").hide();
		stopWaiting();
		return;
	}
	if (data == "QuestionNotExists"){
		ccAlert("QuestionNotExists");
		$("#content-container").hide();
		window.location = "/";
		return;
	}
	
	if (data == "SystemError"){
		ccAlert("SystemError");
		$("#content-container").hide();
		return;
	}
	translate();
	
    viewQuestionModel = ko.mapping.fromJS(data);
    //ko.options.useOnlyNativeEvents = true;
    ko.applyBindings(viewQuestionModel, document.getElementById("question-container"));
    

    //bind Tags
    var tags = data.Tags.trimStr(",").split(",");
    var tagHtml = "";
    
    for (var i = 0; i < tags.length; i++) {
        var tagNameEncode = encodeURI(tags[i]);
        tagHtml += "<a class='btn btn-default btn-xs' href='/tags/tagquestion?title=" + tagNameEncode + "'>"+ tags[i]+"</a>";
    }
    $("#addTags").after(tagHtml);
	
	
    var aid = getQueryString("aid");
    Answers.bindAnswers(data.Id, "answers-container", aid);
    
    translate();
    stopWaiting();
    //CKEDITOR.replace("myAnswer");
}

function getRouteId(){
	var url = window.location.href
	if (url.indexOf("?") > -1){
		url = url.substring(0, url.indexOf("?"))
	}
	var id = url.substring(url.lastIndexOf("/") + 1)
	return id
}
function submitAnswer() {
    //var id = getQueryString("qid");
	var id = getRouteId()

    if (CKEDITOR.instances.myAnswer.getData() == "") {
        ccAlert("AnswerContentIsNull");
        return false;
    }

    var answer = {
        QuestionID: id,
        Content: CKEDITOR.instances.myAnswer.getData()
    }

    Answers.addAnswer(answer, function (msg) {
    	if(msg == "Success"){
    		ccAlert("AddAnswerSuccess");
    		CKEDITOR.instances.myAnswer.setData(''); 
    	}
    	else{
    		ccAlert(msg);
    	}
    });

    return true;
}
function favoriteQuestion(id) {
    Questions.favoriteQuestion(id, function(msg){
    	if(msg == "Success"){
    		ccAlert("FavoriteQuestionSuccess");
    	}
    	else{
    		ccAlert(msg);
    	}

    });
}
function voteAnswer(id, vm) {
    var qid = getRouteId()//getQueryString("qid");
    Answers.voteAnswer(qid, id, function (msg) {
    	if (msg == "Success"){
    	    ccAlert("VoteAnswerSuccess");
        	vm.VotedCount(vm.VotedCount() + 1);
        }
        else {
        	ccAlert(msg);
        }
    });
}
function favoriteAnswer(id) {
    //var qid = getQueryString("qid");
	var qid = getRouteId()
    Answers.favoriteAnswer(qid, id, function (msg) {
    	if (msg == "Success"){
        	ccAlert("FavoriteAnswerSuccess");
        }
        else {
        	ccAlert(msg);
        }
	});
}
function bestAnswer(id,parent, vm) {
    /*
    if (viewQuestionModel.BestAnswer.AnswerContent() !== '') {
        ccAlert('OnlyOneBest');
        return
    }
	*/
	//var qid = getQueryString("qid");
	var qid = getRouteId()
    Answers.bestAnswer(qid,id, function (msg) {
        if (msg == "Success"){
	        ccAlert("BestAnswerSuccess");
	        ko.utils.arrayForEach(parent.items(), function(item) {
		        if(item.IsBest()){
		        	item.IsBest(false);
		        }
		    });

	        vm.IsBest(true);
	        viewQuestionModel.BestAnswer.UserName(vm.UserName());
	        viewQuestionModel.BestAnswer.CreateDate(vm.CreateDate());
	        viewQuestionModel.BestAnswer.AnswerContent(vm.AnswerContent());
	        viewQuestionModel.BestAnswer.DisplayName(vm.DisplayName());
	        viewQuestionModel.BestAnswer.AuthorEmail(vm.AuthorEmail());
	
	        bindUserPresence();
	    	//ProcessImn();
	
	        //$(".btnBest").hide();
	    }
        else{
        	ccAlert(msg);
        }
    });
}
function expertAnswer(id, parent, vm) {
	/*
    if (viewQuestionModel.ExpertAnswer.AnswerContent() !== '') {
        ccAlert('OnlyOneExpert');
        return
    }
	*/
	
	//var qid = getQueryString("qid");
	var qid = getRouteId()
    Answers.expertAnswer(qid,id,function (msg) {
    	if (msg == "Success"){
	        ccAlert("ExpertAnswerSuccess");
	        
	        ko.utils.arrayForEach(parent.items(), function(item) {
		        if(item.IsExpert()){
		        	item.IsExpert(false);
		        }
		    });
		    
	        vm.IsExpert(true);
	        viewQuestionModel.ExpertAnswer.UserName(vm.UserName());
	        viewQuestionModel.ExpertAnswer.CreateDate(vm.CreateDate());
	        viewQuestionModel.ExpertAnswer.AnswerContent(vm.AnswerContent());
	        viewQuestionModel.ExpertAnswer.DisplayName(vm.DisplayName());
	        viewQuestionModel.ExpertAnswer.AuthorEmail(vm.AuthorEmail());
			
	        bindUserPresence();
	    	//ProcessImn();
	    	
	        //$(".btnExpert").hide();
	    }
	    else{
	    	ccAlert(dict[msg]);
	    }
    });
}
function shareQuestion() {
    $("#shareUrl").val(location.href);
    $("#dlgShare").modal('show');
}

function followUser(userName,isFollowed, vm) {
    Users.followUser(userName,isFollowed, function (msg) {
		ccAlert(msg);
		
		if(msg.toLowerCase().indexOf("success") > -1 || msg == "HadFollowedUser" ||  msg == "NotFollowing"){
			vm.IsFollowed(!isFollowed);
			$.each(Answers.answerViewModel.items(), function (it) {
				if(this.UserName() == userName){
					this.IsFollowed(!isFollowed);
				}
	        });
			if(viewQuestionModel.UserName() == userName){
				viewQuestionModel.IsFollowed(!isFollowed);
			}
		}
    });
}

function deleteQuestion(id){
	if(confirm(dict["ConfirmDeleteQuestion"])){
		Questions.deleteQuestion(id, function(msg){
			if(msg == "Success"){
				ccAlert("DeleteQuestionSuccess");
				location.href = "/";
			}
			else{
				ccAlert(msg);
			}
		});
	}
}

