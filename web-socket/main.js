const express = require('express');
const http = require('http');
const socketIO = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketIO(server, {
    cors: {
      origin: "http://localhost:5173", // İstemci adresini buraya yazın
      methods: ["GET", "POST"]
    }
  });

const port = 3008;

const rooms = {};
const usernames = {};

io.on('connection', (socket) => {
  console.log('Kullanıcı Giriş yaptı');

  socket.on('disconnect', () => {
    const username = usernames[socket.id];
    console.log(`${username} Çıkış yaptı`);
    delete rooms[socket.id];
    delete usernames[socket.id];
  });

  socket.on('join', (room, username) => {
    rooms[socket.id] = room;
    usernames[socket.id] = username;

    socket.leaveAll();
    socket.join(room);
    console.log(`${username} kullanıcısı ${room} url'sine sahip odaya girdi`);
    socket.emit('join');
  });

  socket.on('send', (msg) => {
    io.in(rooms[socket.id]).emit('recieve', `${usernames[socket.id]} : ${msg}`);
    console.log(rooms);
    console.log(usernames);
  });

  socket.on('recieve', function (message) {
    socket.emit('recieve', message);
  });
});

server.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});