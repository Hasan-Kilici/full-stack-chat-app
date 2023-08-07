<script>
    let roomName = ''

    function createRoom(){
        fetch("http://localhost:3000/create/room", {
            method:"POST",
            mode:"cors",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: 'include',
            body:  JSON.stringify({ name:roomName }),
        }).then(async(data)=>{
            if(data.ok){
                let res = await data.json()
                let token = res.Token;
                window.location.href = `/room/${token}`
            }
        })
    }
</script>
<div>
    <input type="text" bind:value={roomName}>
    <button on:click={createRoom}>Oluştur</button>
</div>