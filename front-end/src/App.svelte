<script>
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./pages/Home.svelte";
  import Register from "./pages/Register.svelte";
  import Login from "./pages/Login.svelte";
  import Create from "./pages/Create.svelte"
  import Chat from "./pages/Chat.svelte"
  
  export let url = "/";
  import Cookies from "js-cookie";
  let User = false;
  let user = Cookies.get("Token");
  async function GetUser(){
  url = window.location.pathname;
  if(user){
    const response = await fetch(`http://localhost:3000/api/user/${user}`, {
        headers: {
          "Content-Type": "application/json",
        }
      });
  
      if (response.ok) {
        const data = await response.json();
        User = data.data;
        console.log(data)
      } else if (response.status === 404) {
        console.log("Kullanıcı bulunamadı.");
      }
    } 
  }
  GetUser();
  
  </script>
  
  
  <Router {url}>
  <nav>
    <Link to="/">Anasayfa</Link>
    {#if User}
      <Link to="/create/room">Oda oluştur</Link>
      {User.Name}
      <Route path="/create/room" component={Create} />
      <Route path="/chat/:roomId" let:params >
        <Chat roomId={params.roomId} user={User}/>
    </Route>
    {:else}
    <Link to="/register">Kayıt ol</Link>
    <Link to="/login">Giriş yap</Link>
    <Route path="/register" component={Register} />
    <Route path="/login" component={Login} />
    {/if}
  </nav>
  <Route path="/" component={Home} />
  </Router>