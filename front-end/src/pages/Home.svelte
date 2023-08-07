<script>
let rooms = []

async function ListRooms(){
    fetch("http://localhost:3000/api/rooms").then(async(data)=>{
        let res = await data.json();
        let roomdata = res.data;
        rooms = roomdata;
    })
}

ListRooms()
</script>
<main>
    {#each rooms as room}
        <a href="/room/${room.Token}">{room.Name}</a>
    {/each}
</main>