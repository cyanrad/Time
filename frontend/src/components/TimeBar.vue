<template>
  <div class="h-full w-72 bg-stone-50">
    <PeriodBlock
      v-for="(period, index) in periods"
      :h="computeHeight(period.Duration)"
      :text="period.Title"
      :type="period.Type"
      :get-type-color="getTypeColor"
      @click="$emit('periodPressed', index)"
    />
  </div>
  <div class="mr-14 h-full">
    <TimeArrow
      v-for="period in periods"
      :displace="nextArrow(period.Duration)"
      :h="computeHeight(period.Duration)"
      :type="period.Type"
      :time="Math.floor(getSum() / 60) + ':' + (getSum() % 60)"
      :get-type-color="getTypeColor"
    />
  </div>
</template>

<script lang="ts">
import PeriodBlock from "./TimeBarPeriod.vue";
import TimeArrow from "./TimeArrow.vue";
import type { Period } from "@/components/API_Period";

let sum = 0;

export default {
  components: {
    PeriodBlock,
    TimeArrow,
  },
  props: {
    periods: { type: Array<Period>, required: true },
    getTypeColor: { type: Function, required: true }, // passed to TimeArrow & PeriodBlock
  },
  beforeUpdate() {
    sum = 0;
  },
  emits: ["periodPressed"],
  methods: {
    computeHeight(duration: number): number {
      let height: number = window.innerHeight;
      const minutes: number = 1440;
      let result = Math.round((duration / minutes) * height);
      return result != undefined ? result : 0;
    },
    nextArrow(duration: number): number {
      sum = sum + duration;
      return this.computeHeight(sum);
    },
    getSum(): number {
      return sum;
    },
  },
};
</script>
