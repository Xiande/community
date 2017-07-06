/* copyright: PACTERA Beijing co. Ltd */
/// <reference path="jquery-1.7.1.min.js" />
(function ($) {
    $.fn.maxlength = function (settings) {
        if (typeof settings == 'number') {
            settings = { maxLen: settings };
        }
        settings = $.extend({}, $.fn.maxlength.defaults, settings);

        function length(el) {
            var parts = el.val(); //alert(parts);
            if (settings.words) { parts = parts.length ? parts.split(/\s+/) : { length: 0 }; }
            return parts.length;
        }

        return this.each(function () {
            var field = this,
        	        $field = $(field),
                    $showSpan = $("<span class='txtInputLength'></span>"),
                    limit = settings.maxLen;
            $showSpan.insertAfter($field);

            ///禁用鼠标右键菜单
            ///$field.bind("contextmenu", function () { return false; });

            function limitCheck(event) {
                ///禁用Ctrl+v粘贴功能
                ///if (event.ctrlKey && event.keyCode == 86) {
                ///    event.keyCode = 0;
                ///    event.returnValue = false;
                ///    return false;
                ///}

                var len = length($field),
        	                exceeded = len >= limit,
        		            code = event.keyCode;
                if (!exceeded)
                    return;

                switch (code) {
                    case 8:  // allow delete
                    case 9:
                    case 17:
                    case 36: // and cursor keys
                    case 35:
                    case 37:
                    case 38:
                    case 39:
                    case 40:
                    case 46:
                    case 65:
                        return;

                    default:
                        return settings.words && code != 32 && code != 13 && len == limit;
                }
            }

            var updateCount = function () {
                var len = length($field),
            	        diff = limit - len,
                        showText = "(" + len + "/" + limit + ")";
                $showSpan.html(showText);

                if (settings.hardLimit && diff < 0) {
                    field.value = settings.words ?
                    // split by white space, capturing it in the result, then glue them back
            		        field.value.split(/(\s+)/, (limit * 2) - 1).join('') :
            		        field.value.substr(0, limit);

                    updateCount();
                }
            };

            if (settings.hardLimit) { $field.keydown(limitCheck); }
            $field.keyup(updateCount).change(updateCount);

            updateCount();
        });
    };

    $.fn.maxlength.defaults = {
        maxLen: 2,
        hardLimit: true,
        words: false
    };

})(jQuery);