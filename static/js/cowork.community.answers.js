window.Answers = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/ajax/answer",
    answerViewModel: {},
    addAnswer: function (answer, callback) {
        $.ajax({
            async: false,
            type: "POST",
            cache: false,
            url: this.serverUrl,
            contentType: "application/json; charset=utf-8",
            data: encodeURI(JSON.stringify(answer)),
            success: function (msg) {
                callback(msg);
                //window.location = CCConfig.L_Menu_BaseUrl;
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("AddAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    bindAnswers: function (questionId, targetId, aId) {
		//alert(questionId)
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);

            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    url: Answers.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "bindanswers",
                        pageIndex: page,
                        pageSize: self.pageSize,
                        qid: questionId,
                        aid: aId,
                        sortField: "CreateDate",
                        sortOrder: "Desc"
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
                        self.items.removeAll();
						if (data.result != null){
	                       	ko.utils.arrayForEach(data.result, function (answer) {
	                            var t = {
	                                Id: ko.observable(answer.Id),
	                                UserName: ko.observable(answer.UserName),
	                                DisplayName: ko.observable(answer.DisplayName),
	                                IsBest: ko.observable(answer.IsBest),
	                                BestTime: ko.observable(answer.BestTime),
	                                IsExpert: ko.observable(answer.IsExpert),
	                                ExpertTime: ko.observable(answer.ExpertTime),
	                                VotedCount: ko.observable(answer.VotedCount),
	                                CreateDate: ko.observable(answer.CreateDate),
	                                CreateBy: ko.observable(answer.CreateBy),
	                                AnswerContent: ko.observable(answer.AnswerContent),
	                                CanBest: ko.observable(answer.CanBest),
	                                CanExpert: ko.observable(answer.CanExpert),
	                                AuthorEmail: ko.observable(answer.AuthorEmail),
									PhotoImgSrc: ko.observable(answer.PhotoImgSrc),
	                                IsAuthor: ko.observable(answer.IsAuthor),
	                                IsAdmin: ko.observable(answer.IsAdmin),
	                                IsModerator: ko.observable(answer.IsModerator),
	                                IsFollowed: ko.observable(answer.IsFollowed)
	                            			};
	                            self.items.push(t);
	                        });	
						}
                        self.loading(false);
                        bindUserPresence();
                        //ProcessImn();
                        
                        CKEDITOR.replace('myAnswer');
                        $("#loading").hide();
                    },
                    error:function(res){
                    	$("#loading").hide();
                    }
                });
            }]);
        }
        Answers.answerViewModel = new viewModel();
        ko.applyBindings(Answers.answerViewModel, $("#" + targetId)[0]);
        Answers.answerViewModel.goToFirst();

        $("#btnSubmit").click(function () {
            Answers.answerViewModel.goToLast();
        });
    },
    voteAnswer: function (qId, aId, callback) {
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "voteanswer",
                qid: qId,
                aid: aId
            },
            success: function (msg) {
                if (callback != undefined) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("VoteAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    bestAnswer: function (qId, aId, callback) {
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "checkstatus",
                qid: qId,
                aid: aId,
                status: "best"
            },
            success: function (msg) {
                if (msg == "SystemError"){
                    ccAlert("SystemError");
                    return;
                }
            	
                if (msg == "False") {
                    if (confirm(dict["HadBest"])){
                        Answers.best(qId, aId, callback);
                    }
                }
                else if(msg == "True"){
                	Answers.best(qId, aId, callback);
                }
                else{
                	ccAlert(dict[msg]);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("BestAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    best:function(qId, aId, callback){
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "bestanswer",
                qid: qId,
                aid: aId
            },
            success: function (msg) {
                if (callback != undefined) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("BestAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    expertAnswer: function (qId, aId, callback) {
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "checkstatus",
                qid: qId,
                aid: aId,
                status: "expert"
            },
            success: function (msg) {
                if (msg == "SystemError"){
                    ccAlert("SystemError");
                    return;
                }
            	
                if (msg == "False") {
                    if (confirm(dict["HadExpert"])){
                        Answers.expert(qId, aId, callback);
                    }
                }
                else if(msg == "True"){
                	Answers.expert(qId, aId, callback);
                }
                else{
                	ccAlert(dict[msg]);
                }

            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("ExpertAnswerError", XMLHttpRequest.statusText);
            }
        });

    },
    expert: function (qId, aId, callback) {
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "expertanswer",
                qid: qId,
                aid: aId
            },
            success: function (msg) {
                if (callback != undefined) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("ExpertAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    favoriteAnswer: function (qId, aId, callback) {
        $.ajax({
            url: Answers.serverUrl,
            cache: false,
            data: {
                op: "favoriteanswer",
                qid: qId,
                aid: aId
            },
            success: function (msg) {

                if (callback != undefined) {
                    callback(msg);
                }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("FavoriteAnswerError", XMLHttpRequest.statusText);
            }
        });
    },
    GetQuesByVoteAnswer: function (targetID) {
        var VoetAnswerModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    type: "get",
                    url: Answers.serverUrl,
                    cache: false,
                    dataType: "json",
                    contentType: "application/x-www-form-urlencoded; charset=UTF-8",
                    data: {
                        op: "getvoteques",
                        pIndex: page,
                        pSize: self.pageSize
                    },
                    success: function (data) {
                        if (data != null && data != "") {
                            pageHandler.call(self, data);
                            self.items(data.result);
                            bindUserPresence();
                            //ProcessImn();

                        }
                        self.loading(false);
                    },
                    error: function (res) {
                        debugger;
                    }
                });
            }]);
        }
        var voteAnswer = new VoetAnswerModel();
        ko.applyBindings(voteAnswer, $("#" + targetID)[0]);
        voteAnswer.goToFirst();
    },
    GetQuesByFavAnswer: function (targetID) {
        var FavAnswerModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    type: "get",
                    url: Answers.serverUrl,
                    cache: false,
                    dataType: "json",
                    contentType: "application/x-www-form-urlencoded; charset=UTF-8",
                    data: {
                        op: "getfavanswer",
                        pIndex: page,
                        pSize: self.pageSize
                    },
                    success: function (data) {
                        if (data != null && data != "") {
                            pageHandler.call(self, data);
                            self.items(data.result);
                            //bindUserPresence();
                            //ProcessImn();

                        }
                        self.loading(false);
                    },
                    error: function (res) {
                        debugger;
                    }
                });
            }]);
        }
        var favAnswer = new FavAnswerModel();
        ko.applyBindings(favAnswer, $("#" + targetID)[0]);
        favAnswer.goToFirst();
    },
    StopfollowAnswer: function (aid, parent, data) {
        $.ajax({
            type: "get",
            url: Answers.serverUrl,
            cache: false,
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            data: {
                op: "removefa",
                faid: aid
            },
            success: function (res) {
                if (res == "success") {
                    alert(dict["RemoveFavAnswerSuccess"]);
                    parent.items.remove(data);
                } else {
                    alert(dict["RemoveFavAnswerError"]);
                }
            },
            error: function (res) {
                alert(dict["RemoveFavAnswerError"]);
                debugger;
            }
        });
    }
}

