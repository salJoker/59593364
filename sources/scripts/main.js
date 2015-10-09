requirejs.config({
    baseUrl : "js/",

    paths: {
        jquery:'jquery/jquery-1.11.3.min',
        json2:"json2",
        bootstrap:"bootstrap",
        npm:"npm"
    },
    shim:{
        'backbone':{
            deps:['underscore'],
            exports:'Backbone'
        }
    }
});

requirejs(['jquery','json2','backbone','underscore'],function($,JSON){
});

requirejs(['bootstrap','npm'],function(){
});