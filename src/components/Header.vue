<template>
  <div>
    <v-app-bar color="deep-purple-darken-3" app>
      <v-spacer />

      <v-btn class="mx-2" ripple small to="/"><span>Home</span></v-btn>
      <v-btn class="mx-2" ripple small to="/sources"><span>Sources</span></v-btn>
      <v-btn class="mx-2" ripple small to="/destinations"><span>Destinations</span></v-btn>
      <v-btn class="mx-2" ripple small to="/config"><span>Config</span></v-btn>

      <v-spacer />
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {}
};
</script>

<style></style>
