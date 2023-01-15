<template>
  <div>
    <HeaderComponent />
    <!-- <SideBar/> -->
    <router-view />
    <FooterComponent />
  </div>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
// import SideBar from "./components/SideBar.vue";
import FooterComponent from "./components/FooterComponent.vue";


export default {
  name: "PaperManagement",
  components: {
    HeaderComponent,
    // SideBar,
    FooterComponent,
  },
  data() {
    return {
    }
  },
  methods: {
  },
}
</script>

<style scoped>
#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
