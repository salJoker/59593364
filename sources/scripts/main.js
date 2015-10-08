requirejs.config({
    baseUrl : "js/",

    paths: {
        jquery:'jquery/jquery-1.11.3.min'
    }
});

requirejs(['jquery','json2'],function($,JSON){
});