// ==UserScript==
// @name         Trivia
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       You
// @match        https://jackbox.tv/
// @grant        none
// ==/UserScript==

(function() {
    'use strict';

    var delay = 100;

    function eventFire(el, etype)
    {
        if(el.fireEvent)
        {
            el.fireEvent('on' + etype);
        }

        else
        {
            var evObj = document.createEvent('Events');
            evObj.initEvent(etype, true, false);
            el.dispatchEvent(evObj);
        };
    };

    window.fast = function()
    {
        setTimeout(function()
                   {
            var CHILDREN = document.getElementById("prompt").children[0].children[0].children[0].innerText

            var mult = document.getElementsByClassName("choice-button");
            for(var i = 0; i < mult.length; i++) {
                if(mult[i].innerText == eval(CHILDREN)) {
                    eventFire(mult[i], "click")
                }
            }

            window.fast();
        }, delay)
    };
})();

