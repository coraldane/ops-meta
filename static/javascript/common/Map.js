Array.prototype.remove = function (s) {
	for (var i = 0; i < this.length; i++) {
		if (s == this[i]) {
			this.splice(i, 1);
		}
	}
};
/**  
 * Simple Map
 *   
 * var m = new Map();  
 * m.put('key','value');  
 * ...  
 * var s = "";  
 * m.each(function(key,value,index){  
 *      s += index+":"+ key+"="+value+"n";  
 * });  
 * alert(s);  
 *   
 * @author dewitt  
 * @date 2008-05-24  
 */
function Map() {
	/** Key Array */
	this.keys = new Array();
	this.data = new Object();
	/**  
     * @param {String} key  
     * @param {Object} value  
     */
	this.put = function (key, value) {
		if (this.data[key] == null) {
			this.keys.push(key);
		}
		this.data[key] = value;
	};
	/**  
     * @param {String} key
     * @return {Object} value
     */
	this.get = function (key) {
		return this.data[key];
	};
	/**  
     * @param {String} index
     * @return {Object} value
     */
	this.find = function (index) {
		var entrys = this.entrys();
		return this.get(entrys[index].key);
	};
	/**  
     * @param {String} key  
     */
	this.remove = function (key) {
		this.keys.remove(key);
		this.data[key] = null;
	};
	/**  
     * For Each 
     * @param {Function} callback function(key,value,index){..}  
     */
	this.each = function (fn) {
		if (typeof fn != "function") {
			return;
		}
		var len = this.keys.length;
		for (var i = 0; i < len; i++) {
			var k = this.keys[i];
			fn(k, this.data[k], i);
		}
	};
	
	this.keySet = function(){
		return this.keys;
	};
	/**  
     * get an array store key,value set(like entrySet())  
     * @return 
     */
	this.entrys = function () {
		var len = this.keys.length;
		var entrys = new Array(len);
		for (var i = 0; i < len; i++) {
			entrys[i] = {key:this.keys[i], value:this.data[i]};
		}
		return entrys;
	};
	
	this.contains = function(key){
		var len = this.keys.length;
		for(var i=0; i < len; i++){
			if(key == this.keys[i]){
				return true;
			}
		}
		return false;
	}
	
	this.isEmpty = function () {
		return this.keys.length == 0;
	};
	/**  
     * return size  
     */
	this.size = function () {
		return this.keys.length;
	};
	/**  
     * override toString   
     */
	this.toString = function () {
		var s = "{",len=this.keys.length;
		for (var i = 0; i < len; i+=1) {
			var k = this.keys[i];
			var v = this.data[k];
			if("string" == typeof(v)){
				s += '"'+k + '":"' + v + '"';
			}else{
				s += '"'+k + '":' + v + '';
			}
			if(i < len-1){s+=",";}
		}
		s += "}";
		return s;
	};
	/**
    *Clear all elements
    */
	this.clear = function () {
		this.keys = new Array();
		this.data = new Object();
	};
}

