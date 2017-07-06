window.Tags = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/ajax/tags",
    recommendTags: function (title) {
        $.ajax({
            type: "Get",
            cache: false,
            dataType: "json",
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            url: this.serverUrl + "?op=recommendtag&title=" + encodeURI(title),
            success: function (tagList) {
                if (tagList == null) {
                    return;
                }
                var existTag = [];
                $(".checkedItem").each(function (index, element) {
			        var tag = $(element).children(".tagItem").eq(0).text().toLowerCase();
					existTag.push(tag);			        
			    });
                
                for (var i = 0; i < tagList.length; i++) {
                	if(existTag.indexOf(tagList[i].Name) == -1){
                    	addTag(tagList[i].Name);
                    }
                }
            },
            error: function (err) {
                debugger;
            }
        });
    },
    getTagQuestionPage: function (targetID) {
    	var langfilter = getCookie("LangFilter");
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            self.TagTitle = ko.observable(GetRequestPara("title"));
            self.TagID = ko.observable("");            
            self.IsMyFavorite = ko.observable(false);
            self.SelectedLang = ko.observable(langfilter == null ? "all" : langfilter);            
            self.LangList = ko.observableArray([
	            {"langValue": "all","langName": dict["FilterAll"]},
	            {"langValue": "en","langName": dict["FilterEN"]},
	            {"langValue": "cn","langName": dict["FilterCN"]}]);
	        self.langChanged = function(){
	        	if(self.SelectedLang() != undefined && self.SelectedLang() != null){
					self.goToFirst();
					setCookie("LangFilter",self.SelectedLang(), 'd30');
	        	}
	        };
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    type: "get",
                    url: Tags.serverUrl,
                    cache: false,
                    dataType: "json",
                    contentType: "application/x-www-form-urlencoded; charset=UTF-8",
                    data: {
                        op: "tagques",
                        title: self.TagTitle,
                        pIndex: page,
                        pSize: self.pageSize,
                        keyOne: self.KeyOne,
                        keyTwo: self.KeyTwo,
                        sort : self.SortField,
                        lang: self.SelectedLang
                    },
                    success: function (data) {
                        if (data != null && data != "") {
                            pageHandler.call(self, data);
                            self.items(data.result);
                            self.TagID(data.TagID);
                            self.IsMyFavorite(data.IsMyFavorite);
                            //bindUserPresence();
                        	//ProcessImn();

                        }
                        self.loading(false);
                    },
                    error: function (res) {
                    
                    }
                });
            }]);
            self.Sort = function (sortField, data) {
                self.SortField = sortField;
                self.goToFirst();
                return true;
            };

        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();

		$("#refTagQuestion").bind("click", function () {
			vm.KeyOne = $("#tkeyOne").attr("value");
            vm.KeyTwo = "";
            vm.goToFirst();
        });

		
        $("#tsearchFirst").bind("click", function () {
            vm.KeyOne = $("#tkeyOne").attr("value");
            vm.KeyTwo = "";
            vm.goToFirst();
        });
    },
    getFollowedTagsPage: function (targetID) {
        var followedModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $("#layerloading").css("display", "block");
                $.ajax({
                    type: "get",
                    url: Tags.serverUrl,
                    cache: false,
                    dataType: "json",
                    contentType: "application/x-www-form-urlencoded; charset=UTF-8",
                    data: {
                        op: "followedtag",
                        pIndex: page,
                        pSize: self.pageSize
                    },
                    success: function (data) {
                        if (data != null && data != "") {
                            pageHandler.call(self, data);
                            self.items(data.result);
                        }
                        self.loading(false);
                        $("#layerloading").css("display", "none");
                    },
                    error: function (res) {
                        $("#layerloading").attr("display", "none");
                    }
                });
            }]);
        }
        var follow = new followedModel();
        ko.applyBindings(follow, $("#" + targetID)[0]);
        follow.goToFirst();
    },
    SetFavoriteTag: function (tagID, data) {
        $.ajax({
            type: "Get",
            cache: false,
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            url: Tags.serverUrl + "?op=followtag&tagID=" + tagID,
            success: function (result) {
                data.IsMyFavorite(true);
               	ccAlert(result);

            },
            error: function (err) {
                ccAlert("FollowTagError");
                debugger;
            }
        });
    },
    RemoveFavoriteTag: function (favtagID, parent, data) {
        $.ajax({
            type: "Get",
            cache: false,
            contentType: "application/x-www-form-urlencoded; charset=UTF-8",
            url: Tags.serverUrl + "?op=removetag&favtagID=" + favtagID,
            success: function (result,callback) {
                ccAlert("RemoveFavTagSuccess");
                parent.items.remove(data);
            },
            error: function (err) {
                ccAlert("RemoveFavTagError");
                debugger;
            }
        });
    },
    GetTagsPage: function (targetID,size) {
        var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.KeyOne = ko.observable("");
            self.items = ko.observableArray([]);
            //self.items = [];
            paginationViewModel.apply(self, [size, function (page, pageHandler) {
                self.loading(true);
                $("#layerloading").css("display", "block");
                $.ajax({
                    type: "get",
                    url: Tags.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        op: "gettagbyname",
                        tName: self.KeyOne,
                        pIndex: page,
                        pSize: self.pageSize
                    },
                    success: function (data) {
                        if (data != null && data != "") {
                            pageHandler.call(self, data);
                            self.items.removeAll();
							if (data.result != null)
	                            ko.utils.arrayForEach(data.result, function (tag) {
	                                var t = {
	                                    ID: ko.observable(tag.Id),
	                                    TagName: ko.observable(tag.TagName),
	                                    IsMyFavorite: ko.observable(tag.IsMyFavorite),
	                                    TagQuestionCount: ko.observable(tag.TagQuestionCount)
	                                };
	                                self.items.push(t);
	                            });
                        }
                        self.loading(false);
                        $("#layerloading").css("display", "none");
                    },
                    error: function (res) {
                        $("#layerloading").attr("display", "none");
                    }
                });
            }]);
        }
		
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetID)[0]);
        vm.goToFirst();

		$("#refTagpage").bind("click", function () {
            vm.KeyOne = $("#tkeyOne").attr("value");
            vm.goToFirst();
        });

		
        $("#tsearchFirst").bind("click", function () {
            vm.KeyOne = $("#tkeyOne").attr("value");
            vm.goToFirst();
        });
        $("#tkeyOne").bind("keypress", function () {
			if(event.keyCode == 13){
				vm.KeyOne = $("#tkeyOne").attr("value");
            	vm.goToFirst();
			}
        });

    }
}