<template>
  <v-container>
    <v-row>

      <v-col cols="6">
        <v-card shaped>
          <v-card-title>
            <span>disk</span>
          </v-card-title>
          <v-card-subtitle>Space related details.</v-card-subtitle>
          <v-card-text>
            <div v-for="item in destinations" v-bind:key="item.name">
              <router-link :to="item.href">{{ item.config.location }}</router-link>
              -
              <strong>X GB</strong> free of <strong>Y GB</strong>.
            </div>
          </v-card-text>

          <v-divider class="mx-4"></v-divider>

          <v-card-text>
            <div v-for="item in sources" v-bind:key="item.name">
              <router-link :to="item.href" v-if="item.type == 's3'">s3://{{ item.config.bucket }}/{{ item.config.path }}</router-link>
            </div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="6">
        <v-card shaped>
          <v-card-title>
            <span>destinations</span>
          </v-card-title>
          <v-card-subtitle>Manage any Destination.</v-card-subtitle>
        </v-card>
      </v-col>

      <v-col cols="6">
        <v-card shaped>
          <v-card-title>
            <span>sources</span>
          </v-card-title>
          <v-card-subtitle>Manage any Source.</v-card-subtitle>
        </v-card>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    destinations: [],
    sources: [],
  }),
  props: [],
  created() {

    this.$store.dispatch('getDestinations').then(() => {
      this.loadDestinations();
    });

    this.$store.dispatch('getSources').then(() => {
      this.loadSources();
    });

  },
  methods: {
    loadDestinations() {
      this.destinations = this.$store.getters.allDestinations;
    },
    loadSources() {
      this.sources = this.$store.getters.allSources;
    },
    ...mapMutations(['resetSources', 'resetDestinations']),
    ...mapActions(['getSources', 'getDestinations']),
  },
  computed: {
    ...mapGetters(['allSources', 'allDestinations']),
  },
};
</script>

<style></style>
