import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () =>
      import(/* webpackChunkName: "explore" */ "../views/Home.vue")
  },
  {
    path: "/explore",
    name: "Explore",

    component: () =>
      import(/* webpackChunkName: "explore" */ "../views/Explore.vue")
  },
  {
    path: "/aboutme",
    name: "About",

    component: () =>
      import(/* webpackChunkName: "explore" */ "../views/About.vue")
  },
  {
    path: "/apidocs",
    name: "ApiDocs",

    component: () =>
      import(/* webpackChunkName: "explore" */ "../views/ApiDocs.vue")
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
