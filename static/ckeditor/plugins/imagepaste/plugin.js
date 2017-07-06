/*
 * @file image paste plugin for CKEditor
	Feature introduced in: https://bugzilla.mozilla.org/show_bug.cgi?id=490879
	doesn't include images inside HTML (paste from word): https://bugzilla.mozilla.org/show_bug.cgi?id=665341
 * Copyright (C) 2011-13 Alfonso Martínez de Lizarrondo
 *
 * == BEGIN LICENSE ==
 *
 * Licensed under the terms of any of the following licenses at your
 * choice:
 *
 *  - GNU General Public License Version 2 or later (the "GPL")
 *    http://www.gnu.org/licenses/gpl.html
 *
 *  - GNU Lesser General Public License Version 2.1 or later (the "LGPL")
 *    http://www.gnu.org/licenses/lgpl.html
 *
 *  - Mozilla Public License Version 1.1 or later (the "MPL")
 *    http://www.mozilla.org/MPL/MPL-1.1.html
 *
 * == END LICENSE ==
 *
 * version 1.1.1: Added allowedContent settings in case the Advanced tab has been removed from the image dialog
 */

// Handles image pasting in Firefox
CKEDITOR.plugins.add('imagepaste',
{
    init: function (editor) {
        // v 4.1 filters
        if (editor.addFeature) {
            editor.addFeature({
                allowedContent: 'img[!src,id];'
            });
        }
		function getProperty(array, property, regex) {
	        var result = [],
	            i = 0;
	        for (i = 0; i < array.length; i++) {
	            if (regex.test(array[i][property])) {
	                result.push(array[i]);
	            }
	        }
	        return result;
	    }
	    editor.on('keydown',function(e){
	    	alert(e.keyCode);
	    });
        // Paste from clipboard:
        editor.on('paste', function (e) {
        	var html = "";
        	
            if (window.navigator.userAgent.toLowerCase().indexOf("chrome") > -1) {
                var files = e.data.dataTransfer._.files;
                
                var htmlBlobs =  getProperty(files, "type", /image\/png/);
                for(i = 0;i<htmlBlobs.length;i++){
                    var file = htmlBlobs[i];
                    var reader = new FileReader();
                    reader.readAsDataURL(file);
                    reader.onload = function (e) {
                        var base64String = this.result.split("data:image\/png;base64,")[1];
                        var id = CKEDITOR.tools.getNextId();
                        var pData = {
                        	data: encodeURIComponent(base64String),
	                        op:"paste",
	                        ext:".png"
                        };
                        showLoading();
                        $.ajax({//async: false,
                            type: "POST",
                            cache: false,
                            contentType:"application/json; charset=utf-8",
	                        url: editor.config.filebrowserImageUploadUrl ,
	                        data: JSON.stringify(pData),
                            success: function (msg) {
                                var imageUrl = msg;
                                var ele = editor.document.createElement('img',{
                                	attributes:{
                                		src:imageUrl
                                	}
                                });
                                editor.insertElement(ele);
                                $("#loading").hide();
                            },
                            error: function (XMLHttpRequest, textStatus, errorThrown) {
                                alert(XMLHttpRequest.statusText);
                                $("#loading").hide();
                            }
                        });
                    }
                } 
                var data = e.data,
				html = (data.html || (data.type && data.type == 'html' && data.dataValue));
                if (!html){
                    return;
               	}
               	//var fileImage = html.match(/(<img.*src=\"file:\/\/\/.*\"? \/>)/);
               	var fileImage = html.match(/<img.*?(file?:>|\/>)/g);


                if(fileImage != null){
					html = html.replace(/<img.*?(file?:>|\/>)/g,"");
					ccAlert("SinglePasteImageMust");					
					if (e.data.html)
	                    e.data.html = html;
	                else
	                    e.data.dataValue = html;
				}
            }
            else 
            {
                var data = e.data,
				html = (data.html || (data.type && data.type == 'html' && data.dataValue));
                if (!html){
                    return;
                }
               
               	//var fileImage = html.match(/(<img.*src=\"file:\/\/\/.*\"? \/>)/);
               	var fileImage = html.match(/<img.*?(file?:>|\/>)/g);


                if(fileImage != null){
					html = html.replace(/<img.*?(file?:>|\/>)/g,"");
					ccAlert("SinglePasteImageMust");
					if (e.data.html)
	                    e.data.html = html;
	                else
	                    e.data.dataValue = html;
				}
				
				//html.match(/"data:image\/png;base64,(.*?)"/);
                // strip out webkit-fake-url as they are useless:
                if (CKEDITOR.env.webkit && (html.indexOf("webkit-fake-url") > 0)) {
                    alert("Sorry, the images pasted with Safari aren't usable");
                    window.open("https://bugs.webkit.org/show_bug.cgi?id=49141");
                    html = html.replace(/<img src="webkit-fake-url:.*?">/g, "");
                }

                var matchImg = html.match(/"data:image\/png;base64,(.*?)"/);
                if (matchImg == null) {
                    return;
                }

                var base64img = html.match(/"data:image\/png;base64,(.*?)"/)[1];
                html = html.replace(/<img src="data:image\/png;base64,.*?" alt="">/g, function (img) {
                    var data = img.match(/"data:image\/png;base64,(.*?)"/)[1];
                    var id = CKEDITOR.tools.getNextId();
                    var pData = {
                    	data:base64img,
                        op:"paste",
                        ext:".png"
                    };
					showLoading()
                    $.ajax({//async: false,
                        type: "POST",
                        cache: false,
                        contentType:"application/json; charset=utf-8",
                        url: editor.config.filebrowserImageUploadUrl ,
                        data:encodeURIComponent(JSON.stringify(pData)),
                        success: function (msg) {
                            var imageUrl = msg;
                            var theImage = editor.document.getById(id);
                            theImage.data('cke-saved-src', imageUrl);
                            theImage.setAttribute('src', imageUrl);
                            theImage.removeAttribute('id');
                            $("#loading").hide();
                        },
                        error: function (XMLHttpRequest, textStatus, errorThrown) {
                            alert(XMLHttpRequest.statusText);
                            $("#loading").hide();
                        }
                    });

                    return img.replace(/>/, ' id="' + id + '">')
                });
                if (e.data.html)
                    e.data.html = html;
                else
                    e.data.dataValue = html;
            }
        });
    } //Init
});

