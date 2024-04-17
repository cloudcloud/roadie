<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">
                <span>source</span> | <span>{{source_name}}</span>
              </v-col>

              <v-col cols="6">
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headersExpanded" :items="source.entries" :search="search" :single-expand="singleExpand" :expanded.sync="expanded" item-key="entry" show-expand class="elevation-1" v-if="isExpandable">
            <template v-slot:item.action="{ item }">
              <v-btn block small @click="copy(item.entry)">Copy</v-btn>
            </template>
            <template v-slot:expanded-item="{ headers,item }">
              <td :colspan="headers.length">
                <SubSource :sub_name="item.entry" :source_name="source.source.name" />
              </td>
            </template>
          </v-data-table>

          <v-data-table v-else :headers="headers" :items="source.entries" :search="search" class="elevation-1">
            <template v-slot:item.action="{ item }">
              <v-btn block small @click="copy(item.entry)">Copy</v-btn>
            </template>
          </v-data-table>

          <v-dialog v-model="dialog" max-width="500">
            <v-card :loading="loading" class="mx-auto">
              <v-card-title>
                Copy
              </v-card-title>

              <v-card-subtitle>
                Copying {{ entry }}. What is the desired Destination?
              </v-card-subtitle>

              <v-card-text>
                <v-combobox dense outlined persistent-hint solo autofocus v-model="destination" item-text="name" :items="destinations"></v-combobox>
              </v-card-text>

              <v-card-actions>
                <v-spacer />
                <v-btn @click="close">Cancel</v-btn>
                <v-btn @click="save">Save</v-btn>
              </v-card-actions>

            </v-card>
          </v-dialog>

        </v-card>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import SubSource from './SubSource';

export default {
  data: () => ({
    headers: [
      {text: 'Name', align: 'left', value: 'entry'},
      {text: 'Actions', value: 'action'},
    ],
    headersExpanded: [
      {text: 'Expand', value: 'data-table-expand'},
      {text: 'Name', align: 'left', value: 'entry'},
      {text: 'Actions', value: 'action'},
    ],
    source: {},
    dialog: false,
    entry: '',
    loading: false,
    destinations: [],
    destination: '',
    search: '',
    expanded: [],
    singleExpand: true,
    isExpandable: false,
  }),
  props: ['source_name'],
  created() {
    this.loadSource();
  },
  methods: {
    close() {
      this.dialog = false;
    },
    copy(entry) {
      this.entry = entry;
      this.dialog = true;
      this.destinations = this.$store.getters.allDestinations;
    },
    loadSource() {
      this.$store.dispatch('getSource', this.source_name.trim().toLowerCase()).then(() => {
        this.source = this.$store.getters.allSource;
        this.isExpandable = parseInt(this.source.source.config["depth"]) > 0;
      });
    },
    save() {
      this.loading = true;
      this.$store.dispatch('pushCopy', {
        source_name: this.source.source.name,
        entry_name: this.entry,
        destination_name: this.destination.name,
      }).then(() => {
        this.loading = false;
        this.close();
      });
    },
    ...mapMutations(['resetSource']),
    ...mapActions(['getSource']),
  },
  computed: {
    ...mapGetters(['allSource', 'getCopyState']),
  },
  components: {
    SubSource,
  },
};
</script>

<style></style>
