<template>
  <div class="footer">
    <div class="p-grid p-pl-6">
      <div class="p-col">
        <h3>About</h3>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Maiores sed
          doloribus, ea dicta quaerat illo vero ad esse adipisci neque nihil.
          Corrupti quod sequi rerum eveniet enim quasi incidunt porro.
        </p>
      </div>
      <div class="p-col">
        <h3>Quick Links</h3>
        <p>
          <router-link to="/">Home</router-link><br />
          <router-link to="/explore">Explore</router-link><br />
          <router-link to="/apidocs">API Docs</router-link><br />
        </p>
      </div>
    </div>
    <div class="sub-footer">
      <i class="pi pi-github"></i>
      <i class="pi pi-inbox" style="margin-left:2vw"></i>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.footer {
  background-color: var(--footer-background-color);
}
.sub-footer {
  background-color: darkgray;
}
h3 {
  font-size: 1.5em;
  color: var(--footer-h3);
}
i {
  font-size: 2rem;
  margin: 20px 0px 10px 40px;
}
p {
  color: var(--footer-text);
  font-size: 1.2em;
}
a {
  text-decoration: none;
  color: var(--footer-text);
}
a:hover {
  color: black;
  color: var(--footer-link);
}
</style>
