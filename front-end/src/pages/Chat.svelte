<script>
    import { onMount } from 'svelte';
    import { io } from 'socket.io-client';
  
    let socket;
    export let user;
    export let roomId;
    let username = user.Name;
    let message = '';
    let messages = [];
    let users = [];
  
    const socketUrl = 'http://localhost:3008';
  
      socket = io(socketUrl);
  
      socket.on('updateUsers', (updatedUsers) => {
        users = updatedUsers;
      });
  
      socket.on('message', (msg) => {
        messages.push(msg);
        messages = messages
        console.log(msg)
      });

      socket.on('sendMessage', (msg) => {
        messages.push(msg);
        messages = messages
        console.log(msg)
      });
  
    function joinRoom() {
      socket.emit('joinRoom', { roomId, username });
    }
  
    function sendMessage() {
    socket.emit('sendMessage', { roomId, message, username });
    messages.push({ username, message });
    messages = messages;
    message = '';
  }
  </script>
  
  <main>
    <h1>Chat App</h1>
    <div>
      {#each users as user}
        <p>{user.username} - {user.roomId}</p>
      {/each}
    </div>
    <br />
    <input type="text" bind:value="{message}" placeholder="Message" />
    <button on:click="{sendMessage}">Send</button>
    <br />
    <div>
      {#each messages as msg}
        <p><img src="https://ui-avatars.com/api/?size=40&font-size=0.66&name={msg.username}2i&background=random&color=random?format=txt">{msg.username}: {msg.message}</p>
      {/each}
    </div>
  </main>
  