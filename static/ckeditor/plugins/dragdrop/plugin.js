CKEDITOR.plugins.add( 'dragdrop', {
    init: function(editor) {
        backends = {
            imgur: {
                upload: uploadImgur,
                required: ['clientId'],
                init: function() {}
            },
            s3: {
                upload: uploadS3,
                required: [
                    'bucket', 'accessKeyId','secretAccessKey', 'region'
                ],
                init: function() {
                    var script = document.createElement('script');
                    script.async = 1;
                    script.src = 'aws-sdk-2.1.26.js';
                    document.body.appendChild(script);
                }
            }
        };

        var checkRequirement = function(condition, message) {
            if (!condition)
                throw Error("Assert failed" + (typeof message !== "undefined" ? ": " + message : ""));
        };

        function validateConfig() {
            var errorTemplate = 'DragDropUpload Error: ->';
            checkRequirement(
                editor.config.hasOwnProperty('dragdropConfig'),
                "Missing required dragdropConfig in CKEDITOR.config.js"
            );

            var backend = backends[editor.config.dragdropConfig.backend];

            var suppliedKeys = Object.keys(editor.config.dragdropConfig.settings);
            var requiredKeys = backend.required;

            var missing = requiredKeys.filter(function(key) {
                return suppliedKeys.indexOf(key) < 0
            });
            if (missing.length > 0) {
                throw 'Invalid Config: Missing required keys: ' + missing.join(', ')
            }
        }

        validateConfig();

        var backend = backends[editor.config.dragdropConfig.backend];
        backend.init();

        function doNothing(e) { }
        function orPopError(err) { alert(err.data.error) }

        function dropHandler(e) {
        	e.preventDefault();
            var file = e.dataTransfer.files[0];
            var reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = function (e) {
                var base64String = this.result.split("base64,")[1];
                var splitStr = file.name.split('.');
                var ext = "." + splitStr[splitStr.length -1];
                if(".jpg;.gif;.bmp;.png;".indexOf(ext) > -1){
                	backend.upload(ext, base64String).then(insertImage, orPopError);
                }else{
                	ccAlert("FileFormatTip");
                }
                
            }
        }

        function insertImage(href) {
            //var editor = CKEDITOR.instances.editor01;
            var elem = editor.document.createElement('img', {
                attributes: {
                    src: href
                }
            });
            editor.insertElement(elem);
        }

        function uploadImgur(ext,base64img) {
        	 return new Promise(function(resolve, reject) {
				var postData = {
	                	data:base64img,
	                	op:"drop",
	                	ext:ext
	                };
	        	 $.ajax({//async: false,
	                type: "POST",
	                cache: false,
	                contentType:"application/json; charset=utf-8",
	                url: editor.config.filebrowserImageUploadUrl ,
	                data: encodeURIComponent(JSON.stringify(postData)),
	                success: function (msg) {
						resolve(msg);
	                },
	                error: function (XMLHttpRequest, textStatus, errorThrown) {
	                    reject(XMLHttpRequest.responseText);	                    
	                }
	            });
			 });
        }

        function uploadS3(file) {
            var settings = editor.config.dragdropConfig.settings;
            AWS.config.update({accessKeyId: settings.accessKeyId, secretAccessKey: settings.secretAccessKey});
            AWS.config.region = 'us-east-1';

            console.log(settings);
            console.log(AWS.config);

            var bucket = new AWS.S3({params: {Bucket: settings.bucket}});
            var params = {Key: file.name, ContentType: file.type, Body: file, ACL: "public-read"};
            return new Promise(function(resolve, reject) {
                bucket.upload(params, function (err, data) {
                    if (!err) resolve(data.Location);
                    else reject(err);
                });
            });
        };

        CKEDITOR.on('instanceReady', function() {
            var iframeBase = document.querySelector('iframe').contentDocument.querySelector('html');
            var iframeBody = iframeBase.querySelector('body');

            iframeBody.ondragover = doNothing;//doNothing;
            iframeBody.ondrop = dropHandler;
            paddingToCenterBody = ((iframeBase.offsetWidth - iframeBody.offsetWidth) / 2) + 'px';
            iframeBase.style.height = '90%';
            iframeBase.style.width = '90%';
            iframeBase.style.overflowX = 'hidden';

            iframeBody.style.height = '90%';
            iframeBody.style.margin = '0';
            iframeBody.style.paddingLeft = paddingToCenterBody;
            iframeBody.style.paddingRight = paddingToCenterBody;
        });
    }
});
