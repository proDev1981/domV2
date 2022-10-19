package dom


func js() string{
	js:= `
	<script>

	const ws = new WebSocket ("ws://localhost`+ port + rootws +`");

	ws.onopen = (e)=>{

		console.log("conectado");
		ws.send("ok");

	};

	ws.onmessage = (e)=>{
		console.log(e.data);
		let data = JSON.parse( e.data );

		if ( data.type == "bind" ){
			isBind( data );
			return;
		}
		if ( data.type == "eval" ){
			isEval( data );
			return; 
		}
	};

	function isBind( data ){

		window[data.name] = ()=>{
			if ( event.type != "dragstart"){
				event.preventDefault();
			}
			console.log(event.target.value);
			ws.send( JSON.stringify({type:"event", name:data.name ,event:JSON.stringify({type:event.type,ref:event.target.getAttribute('key'),value:event.target.value})}) );
		};
	};

	function isEval( data ){

		let res = eval( data.js );

		if ( typeof res != "string" ){
			res = JSON.stringify( res );
		}
		if ( data.id ){
			res = JSON.stringify( {id : data.id , body: res} );
		}
		if ( res != undefined ){
			ws.send( res );
		}
	};


	window.addEventListener('beforeunload', (e)=>{
		e.preventDefault();
		ws.send("close");
	});

	window.uploadValue = ( data )=>{
		let res = data;
		let ele = document.querySelector('.' + data.ref );
		console.log(ele)

		ele.addEventListener('change',()=>{
				res.value = ele.value.toString();
				ws.send(JSON.stringify({type:"upload", Ref:data.ref , value : res.value , body :JSON.stringify(res)}));
		});
	};
	</script>
	`
	return js
}
