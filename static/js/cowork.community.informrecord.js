window.InformRecord = {
    serverUrl: CCConfig.L_Menu_BaseUrl + "/_layouts/15/Lenovo.Cowork.Community.sp/Ajax/InformRecordOP.ashx",
    addRecord: function (record) {
        $.ajax({
            async: false,
            type: "POST",
            cache: false,
            url: this.serverUrl,
            contentType: "application/json; charset=utf-8",
            data: encodeURI(JSON.stringify(record)),
            success: function (msg) {
            	if (msg == "Success"){
	                ccAlert("AddInformRecordSuccess");
	                var url = GetRequestPara("source");
	                if (url == null) {
	                    url = "index.aspx"
	                }
	
	                window.location = url;
	            }
	            else{
	            	ccAlert(msg);
	            }
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                ccAlert("AddInformRecordError",XMLHttpRequest.statusText);
            }
        });
    },
    getInformList:function(targetId,isMyList){
    	var viewModel = function () {
            var self = this;
            self.loading = ko.observable(true);
            self.items = ko.observableArray([]);
            paginationViewModel.apply(self, [10, function (page, pageHandler) {
                self.loading(true);
                $.ajax({
                    url: InformRecord.serverUrl,
                    cache: false,
                    dataType: "json",
                    data: {
                        OP: "getlist",
                        pageIndex: page,
                        pageSize: self.pageSize,
                        isMyList:isMyList,
                        sortField: "CreateDate",
                       	sortOrder: "Desc"
                    },
                    success: function (data) {
                        pageHandler.call(self, data);
                        self.items(data.result);
                        self.loading(false);
                    }
                });
            }]);
        }
        var vm = new viewModel();
        ko.applyBindings(vm, $("#" + targetId)[0]);
        vm.goToFirst();
    },
    viewInform:function(id,targetId){
    	var url = InformRecord.serverUrl + "?op=viewinform&iid=" + id;
        $.ajax({
            type: "Get",
            url: url,
            cache: false,
            contentType: "application/json; charset=utf-8",
            dataType:"json",
            success: function (data) {
                var vm = ko.mapping.fromJS(data);
				ko.applyBindings(vm, $("#" + targetId)[0]);
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
	            if (XMLHttpRequest.responseText == "InformRecordNotExists"){
					ccAlert("InformRecordNotExists");
					$("#main-panel").hide();
				}
            }
        });
    },
    approveInform:function(id,valid,vm){
    	var url = InformRecord.serverUrl + "?op=approveinform&iid=" + id + "&isvalid=" + valid;
        $.ajax({
            type: "Get",
            url: url,
            cache: false,
            contentType: "application/json; charset=utf-8",
            success: function (msg) {
                if(msg == "Success"){
		    		ccAlert("ApproveSuccess");
		    		var source = getQueryString("source");
		    		window.location = source;
		    	}
		    	else{
		    		ccAlert(msg);
		    	}

            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
	            if (XMLHttpRequest.responseText == "InformRecordNotExists"){
					ccAlert("InformRecordNotExists");
					$("#main-panel").hide();
				}
            }
        });
    }

}
