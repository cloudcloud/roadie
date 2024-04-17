<template>
  <v-container>
    <v-row>

      <v-col cols="6">
        <v-card shaped>
          <v-card-title>
            <span>destinations</span>
          </v-card-title>
          <v-card-subtitle>Manage any Destination.</v-card-subtitle>

          <v-card-text>
            <div v-for="item in config.destinations" v-bind:key="item.name">
              <router-link :to="makeHref('/destinations', item.name)">{{ item.name }} - {{ item.config.location }}</router-link>
              <span v-if="item.disk_info"> :: {{ diskSize(item.disk_info.free) }} free of {{ diskSize(item.disk_info.size) }}.</span>
            </div>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="6">
        <v-card shaped>
          <v-card-title>
            <span>sources</span>
          </v-card-title>
          <v-card-subtitle>Manage any Source.</v-card-subtitle>

          <v-card-text>
            <div v-for="item in config.sources" v-bind:key="item.name">
              <div v-if="item.type == 'local_path'">
                <router-link :to="makeHref('/sources', item.name)">{{ item.name }} - {{ item.config.location }}</router-link>
                <span v-if="item.disk_info"> :: {{ diskSize(item.disk_info.free) }} free of {{ diskSize(item.disk_info.size) }}.</span>
              </div>
              <div v-if="item.type == 's3'">
                <router-link :to="makeHref('/sources', item.name)">{{ item.name }} - s3://{{ item.config.bucket }}/{{ item.config.path }}</router-link>
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    config: [],
  }),
  props: [],
  created() {
    this.$store.dispatch('getConfig').then(() => {
      this.loadConfig();
    });
  },
  methods: {

    diskSize(num) {
      let n = num
      if (n / 1024 > 1) {
        n /= 1024
        if (n / 1024 > 1) {
          n /= 1024
          if (n / 1024 > 1) {
            n /= 1024
            if (n / 1024 > 1) {
              n /= 1024
              if (n / 1024 > 1) {
                n /= 1024
                return n.toLocaleString('en-AU', { style: 'unit', unit: 'petabyte'});
              }
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

    makeHref(path, name) {
      return path + "/" + name;
    },

    loadConfig() {
      this.config = this.$store.getters.allConfig;
    },
    ...mapMutations(['resetConfig']),
    ...mapActions(['getConfig']),
  },
  computed: {
    ...mapGetters(['allConfig']),
  },
};
</script>

<style></style>
