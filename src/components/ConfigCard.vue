<template>
  <v-card shaped>
    <v-card-title><span>{{ title }}</span></v-card-title>
    <v-card-subtitle>{{ subtitle }}</v-card-subtitle>

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
import ConfigEditDialog from './ConfigEditDialog';
import ConfigRemoveDialog from './ConfigRemoveDialog';

export default {
  data: () => ({
  }),
  props: ['title', 'subtitle', 'items', 'headers', 'link-prefix', 'type'],
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
    ConfigEditDialog,
    ConfigRemoveDialog,
  },
};
</script>

<style></style>

