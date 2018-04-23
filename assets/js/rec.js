let Reco=function(){
    let self = this;
	this.rec = undefined;
	this.isStart = false;
	this.isRecoding = false;
    this.endText = "";
    this.recodingText = [];
    this.checkTime = "";
    this.sst = new StringStorage();
	
	this.init = function(){
		console.log('init');
		self.rec = new webkitSpeechRecognition();
		self.rec.continuous = true;
	    self.rec.interimResults = true;
	    self.rec.maxAlternatives = 1;
	    self.rec.lang = 'ja-JP';
		
	};
	this.setup = function(){
	    self.rec.onaudiostart = function(){console.log("rec onaudiostart");}
	    self.rec.onaudioend = function(){console.log("rec onaudioend");}
	    self.rec.onsoundstart = function(){console.log("rec onsoundstart");}
	    self.rec.onsoundend = function(){console.log("rec onsoundend");}
	    self.rec.onspeechstart = function(){console.log("rec onspeechstart");}
	    self.rec.onspeechend = function(){console.log("rec onspeechend");}
	    self.rec.onstart = function(){console.log("rec onstart");}
	    self.rec.onnomatch = function() {console.log('Speech not recognised');}
		self.rec.onend = function(){
			console.log("rec onend");
	        if(self.isStart){
		      console.log("re start");
	          self.rec.start();  
	        }
		}
		self.rec.onerror = function(e){
	        console.log("rec onerror "+e.error);
	        if(!self.isStart){
	          self.rec.stop();
	        }
		}
		self.rec.onresult = function(e) {
	        for (var i = e.resultIndex; i < e.results.length; ++i) {
	            if (e.results[i].isFinal) {
	                var str = e.results[i][0].transcript;
	                console.log('Recognised: ' + str);
                    let localtext = "";
                    console.log('addtext: ' + localtext);
                    self.sst.convert();
                    console.log("SST ENDTEXT : ");
                    console.log(self.sst);
                    
                    var v = $("#area").val();
                    $("#area").val(self.sst.endText);
	            }
	            else{
	                var str = e.results[i][0].transcript;
	                console.log('recoding: ' + str);
                    self.sst.add(str,false);
	            }
	        }
        }
	};
	this.start=function(){
		self.rec.start();
		self.isStart = true;
	};
	this.stop=function(){
		self.rec.stop();
		self.isStart = false;
	};
	this.init();
	this.setup();
    this.checkTime = new Date();
    this.recodingText.push("");
    this.interval = setInterval(function(){
        let now = new Date();
        let time = now.getTime() - self.checkTime.getTime();
        var miri_second = time/100;
        if(miri_second>100){
            if(self.isStart && self.sst.writingTextArray.length>0){
                self.rec.stop();
            }
        }
    },100)
}
    
let StringStorage=function(){
    self=this;
    this.writingTextArray = [];
    this.endText = "";
    this.checkTime = 0;
    this.add=function(text,addArray){
        if(addArray){
            self.writingTextArray.push(text); 

        }
        else{
            let len = self.writingTextArray.length;
            self.writingTextArray[len-1] = text;
        }
    };
    this.slash=function(){
        let len = self.writingTextArray.length;
        let txt = self.writingTextArray[len-1];
        if(txt.slice(-1) != "," && txt.slice(-1) != ""){
            txt += ",";
            self.recodingText[len-2] = txt;
            self.recodingText.push("");
        }
    }
    this.convert=function(){
        let localtext = "";
        for(var i=0;i<self.writingTextArray.length;i++){
            let str = self.writingTextArray[i];
            var regExp = new RegExp( localtext, "y" );
            console.log(regExp);
            let addtext = str.replace(regExp,"");
            console.log("str : "+str+' / addtext :'+addtext+' / localtext :'+localtext);
            if(addtext!=""){
                localtext += addtext;
                self.endText += addtext;
            }
            
        }
        self.endText += "\n";
        self.writingTextArray = []
    }
}
var sst = new StringStorage()
sst.add('試験の準備')
sst.add('試験の準備')
sst.add('試験の準備')
sst.add('試験の準備であると同時の重要な責任を')
sst.add('試験の準備であると同時の重要な責任を果たす')
sst.add('試験の準備であると同時の重要な責任を果たすため')
console.log(sst.writingTextArray);
sst.convert();
console.log(sst.endText);