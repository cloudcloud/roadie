<template>
  <v-card shaped :title="title" :subtitle="subtitle">
    <template v-slot:append>
      <ConfigAddDialog :type="type"></ConfigAddDialog>
    </template>

    <v-card-text>
      <v-data-table-virtual :items="items" :items-per-page="0" :headers="headers" hover no-filter disable-pagination>
        <template v-slot:item.action="{ item }">
          <v-spacer></v-spacer>
          <ConfigEditDialog :name="item.name" :fullObj="item" :type="type"></ConfigEditDialog>
          <ConfigRemoveDialog :name="item.name" :fullObj="item" :type="type"></ConfigRemoveDialog>
          <v-spacer></v-spacer>
        </template>

        <template v-slot:item.name="{ item }">
          <router-link :to="makeHref(item.name)">{{ item.name }}</router-link>
        </template>

        <template v-slot:item.disk_info="{ item }">
          <strong v-if="item.disk_info">{{ diskSize(item.disk_info.free) }}</strong>
        </template>

        <template #bottom></template>
      </v-data-table-virtual>
    </v-card-text>
  </v-card>
</template>

<script>
import ConfigAddDialog from './ConfigAddDialog';
import ConfigEditDialog from './ConfigEditDialog';
import ConfigRemoveDialog from './ConfigRemoveDialog';
import { mdiPlusOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiPlusOutline,
    headers: [
      {title: 'Name', align: 'left', key: 'name'},
      {title: 'Location', align: 'left', key: 'config.location'},
      {title: 'Disk Free', align: 'left', key: 'disk_info'},
      {title: 'Actions', align: 'center', key: 'action', sortable: false},
    ],
  }),
  props: ['title', 'subtitle', 'items', 'link-prefix', 'type'],
  methods: {
    diskSize(num) {
      let n = num;
      let typ = "byte";
      while (n / 1024 > 1) {
        switch (typ) {
          case "byte":
            typ = "kilobyte";
            break;
          case "kilobyte":
            typ = "megabyte";
            break;
          case "megabyte":
            typ = "gigabyte";
            break;
          case "gigabyte":
            typ = "terabyte";
            break;
          case "terabyte":
            typ = "petabyte";
            break;
        }
        n /= 1024;
      }
      return n.toLocaleString('en-AU', { style: 'unit', unit: typ });
    },
    makeHref(val) {
      return this.linkPrefix + val;
    },
  },
  components: {
    ConfigAddDialog,
    ConfigEditDialog,
    ConfigRemoveDialog,
  },
};
</script>

<style></style>

