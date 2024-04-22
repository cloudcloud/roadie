<template>
  <v-container>
    <v-row>

      <v-col cols="6">
        <ConfigCard title="Destinations" subtitle="Manage any Destination" :items="config.destinations" :headers="headers" link-prefix="/destinations/" type="destination"></ConfigCard>
      </v-col>

      <v-col cols="6">
        <ConfigCard title="Sources" subtitle="Manage any Source" :items="config.sources" :headers="headers" link-prefix="/sources/" type="source"></ConfigCard>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import ConfigCard from './ConfigCard';

export default {
  data: () => ({
    config: [],
    headers: [
      {title: 'Name', align: 'left', key: 'name'},
      {title: 'Location', align: 'left', key: 'config.location'},
      {title: 'Disk Free', align: 'left', key: 'disk_info'},
      {title: 'Actions', align: 'center', key: 'action'},
    ],
  }),
  props: [],
  created() {
    this.$store.dispatch('getConfig').then(() => {
      this.loadConfig();
    });
  },
  methods: {
    loadConfig() {
      this.config = this.$store.getters.allConfig;
    },
    ...mapMutations(['resetConfig']),
    ...mapActions(['getConfig']),
  },
  computed: {
    ...mapGetters(['allConfig']),
  },
  components: {
    ConfigCard,
  },
};
</script>

<style></style>
