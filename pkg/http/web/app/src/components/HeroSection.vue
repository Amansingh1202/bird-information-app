<template>
  <div class="home-hero" role="banner">
    <div class="overlay">
      <h1><span class="big-text">Xplore</span> the underwater world!</h1>
      <h2>Search about your favourite fish</h2>
      <button><router-link to="/explore">Explore</router-link></button>
    </div>
    <video class="home-video" autoplay loop muted>
      <source src="../assets/videos/hero.mp4" type="video/mp4" />
    </video>
    <div></div>
  </div>
</template>

<style scoped>
.home-hero {
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 5px;
  background-attachment: scroll;
  overflow: hidden;
}

h1 {
  color: aliceblue;
  font-size: 3em;
  margin-top: 20%;
  margin-left: 30%;
}
h2 {
  color: palevioletred;
  font-size: 2.5em;
  margin-top: -15px;
  margin-left: 34%;
}
span {
  color: cyan;
  font-size: 1.5em;
}

.home-hero video {
  min-width: 100%;
  min-height: 100%;
  position: relative;
  z-index: 1;
}

.home-hero .overlay {
  background-color: rgba(39, 69, 145, 0.44);
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0px;
  bottom: 0px;
  z-index: 2;
}
a {
  color: aqua;
  text-decoration: none;
}
button {
  margin-top: 10px;
  font-size: 1.8em;
  padding: 5px 15px;
  border: none;
  border-radius: 5px;
  background-color: rgb(0, 0, 105);
  margin-left: 46%;
}
button:hover {
  background-color: rgb(2, 2, 248);
  color: aquamarine;
}
</style>
