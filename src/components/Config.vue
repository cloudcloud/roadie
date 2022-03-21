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
            <div v-for="item in condense()" v-bind:key="item.name">
              <div v-if="item.type == 'local_path'">
                <router-link :to="item.href">{{ item.config.location }}</router-link>
                <span v-if="item.hasDisk"> - {{ item.disk.free }} free of {{ item.disk.size }}.</span>
              </div>
              <div v-if="item.type == 's3'">
                <router-link :to="item.href">s3://{{ item.config.bucket }}/{{ item.config.path }}</router-link>
              </div>
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
    everything: [],
    config: [],
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

    this.$store.dispatch('getConfig').then(() => {
      this.loadConfig();
    });

  },
  methods: {
    condense() {
      let everything = [];
      if (this.sources.length < 1 || this.destinations.length < 1 || this.config.length < 1) {
        return everything;
      }

      this.sources.forEach((s) => {
        let done = false;
        this.config.disk_info.forEach((d) => {
          if (d.path == s.config.location) {
            everything.push({
              config: s.config,
              disk: {
                free: this.diskSize(d.free),
                size: this.diskSize(d.size),
              },
              hasDisk: true,
              href: s.href,
              name: s.name,
              type: s.type,
            });
            done = true;
          }
        });

        if (!done) {
          everything.push(s);
        }
      });

      this.destinations.forEach((s) => {
        let done = false;
        this.config.disk_info.forEach((d) => {
          if (d.path == s.config.location) {
            everything.push({
              config: s.config,
              disk: {
                free: this.diskSize(d.free),
                size: this.diskSize(d.size),
              },
              hasDisk: true,
              href: s.href,
              name: s.name,
              type: s.type,
            });
            done = true;
          }
        });

        if (!done) {
          everything.push(s);
        }
      });

      return everything;
    },

    diskSize(num) {
      let n = num
      if (n / 1000 > 1) {
        n /= 1000
        if (n / 1000 > 1) {
          n /= 1000
          if (n / 1000 > 1) {
            n /= 1000
            if (n / 1000 > 1) {
              n /= 1000
              return n.toLocaleString('en-AU', { style: 'unit', unit: 'terabyte'});
            }
            return n.toLocaleString('en-AU', { style: 'unit', unit: 'gigabyte' });
          }
          return n.toLocaleString('en-AU', { style: 'unit', unit: 'megabyte' });
        }
        return n.toLocaleString('en-AU', { style: 'unit', unit: 'kilobyte' });
      }
      return n.toLocaleString('en-AU', { style: 'unit', unit: 'byte' });
    },

    loadDestinations() {
      this.destinations = this.$store.getters.allDestinations;
    },
    loadSources() {
      this.sources = this.$store.getters.allSources;
    },
    loadConfig() {
      this.config = this.$store.getters.allConfig;
    },
    ...mapMutations(['resetSources', 'resetDestinations', 'resetConfig']),
    ...mapActions(['getSources', 'getDestinations', 'getConfig']),
  },
  computed: {
    ...mapGetters(['allSources', 'allDestinations', 'allConfig']),
  },
};
</script>

<style></style>
