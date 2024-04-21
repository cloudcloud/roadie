<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">
                <span>source <v-icon :icon="`${mdiArrowRightThinCircleOutline}`"></v-icon> {{source_name}}</span>
              </v-col>

              <v-col cols="6">
                <v-text-field v-model="search" single-line hide-details clearable placeholder="evil" variant="underlined">
                  <v-icon :icon="`${mdiMagnify}`"></v-icon>
                </v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table-virtual :headers="headersExpanded" :items="source.entries" :search="search" :single-expand="singleExpand" :expanded.sync="expanded" item-value="entry" show-expand class="elevation-1" height="100%" v-if="isExpandable">
            <template v-slot:item.action="{ item }">
              <CopyDialog :source_name="source.source.name" :entry_name="item.entry + '/'" />
            </template>
            <template v-slot:expanded-row="{ columns,item }">
              <tr>
                <td :colspan="columns.length">
                  <SubSource :sub_name="item.entry" :source_name="source.source.name" />
                </td>
              </tr>
            </template>
          </v-data-table-virtual>

          <v-data-table-virtual v-else height="500" :headers="headers" :items="source.entries" :search="search" class="elevation-1" hover>
            <template v-slot:item.action="{ item }">
              <CopyDialog :source_name="source.source.name" :entry_name="item.entry" />
            </template>
          </v-data-table-virtual>

        </v-card>
      </v-col>

    </v-row>
  </v-container>

</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import CopyDialog from './CopyDialog';
import SubSource from './SubSource';
import { mdiArrowRightThinCircleOutline, mdiMagnify } from '@mdi/js';

export default {
  data: () => ({
    mdiArrowRightThinCircleOutline,
    mdiMagnify,
    headers: [
      {title: 'Name', align: 'left', key: 'entry'},
      {title: 'Actions', key: 'action'},
    ],
    headersExpanded: [
      {title: '', key: 'data-table-expand'},
      {title: 'Name', align: 'left', key: 'entry'},
      {title: 'Actions', key: 'action'},
    ],
    source: {},
    entry: '',
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
    loadSource() {
      this.$store.dispatch('getSource', this.source_name.trim().toLowerCase()).then(() => {
        this.source = this.$store.getters.allSource;
        this.isExpandable = parseInt(this.source.source.config["depth"]) > 0;
      });
    },
    ...mapMutations(['resetSource']),
    ...mapActions(['getSource']),
  },
  computed: {
    ...mapGetters(['allSource']),
  },
  components: {
    CopyDialog, SubSource,
  },
};
</script>

<style></style>
