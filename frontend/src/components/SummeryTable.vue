<template>
  <div class="m-4 mb-2 basis-1/2 rounded bg-stone-50 p-2">
    <h1 class="mb-2 text-center font-mono text-lg font-semibold">Today</h1>
    <div class="flex flex-row flex-wrap justify-center">
      <div
        v-for="t in types"
        class="m-1 rounded-lg border-2 p-1.5"
        :style="{
          borderColor: getTypeColor(t, 2),
          backgroundColor: getTypeColor(t, 0),
        }"
      >
        <p class="text-md inline pr-2 font-mono font-semibold">
          {{ t }}
        </p>
        <p
          class="text-md inline font-mono"
          :style="{ color: getTypeColor(t, 2) }"
        >
          {{ calcAmt.get(t) }}
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import type { Period } from "@/components/API_Period.js";
export default {
  props: {
    periods: { type: Array<Period>, required: true },
    types: { type: Array<string>, required: true },
    getTypeColor: { type: Function, required: true },
  },

  computed: {
    // >> Calculates the total amount of time spent on a type of activity
    calcAmt(): Map<string, number> {
      let typeAmt = new Map<string, number>();

      // filling the typeAmt map with the exsisting types
      this.types.forEach((t) => {
        typeAmt.set(t, 0);
      });

      // calculating averages
      this.periods.forEach((p) => {
        let currentVal: number = typeAmt.get(p.Type) ?? 0;
        currentVal += p.Duration;
        typeAmt.set(p.Type, currentVal);
      });

      return typeAmt;
    },
  },
};
</script>
