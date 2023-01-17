<template>
  <!-- post form -->
  <div class="m-4 rounded bg-stone-50" v-if="selectedPeriod.Date_Stamp == 0">
    <div class="p-3">
      <label for="type" class="font-mono text-lg font-semibold">Type</label
      ><br />
      <select
        name="type"
        id="type"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="newPeriod.Type"
      >
        <option v-for="t in types">{{ t }}</option>
      </select>
      <label for="title" class="font-mono text-lg font-semibold">Title</label
      ><br />
      <input
        type="text"
        id="title"
        name="title"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="newPeriod.Title"
      />
      <label for="description" class="font-mono text-lg font-semibold"
        >Description</label
      ><br />
      <textarea
        id="description"
        name="description"
        rows="4"
        class="mb-1 w-full resize-none rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="newPeriod.Description"
      ></textarea>
      <label for="duration" class="font-mono text-lg font-semibold"
        >Duration</label
      ><br />
      <input
        type="text"
        id="duration"
        name="duration"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        :value.number="durationFormula"
        @input="updateDuration"
      />
      <button
        class="w-full justify-center rounded bg-stone-700 text-center font-mono text-xl font-bold text-stone-50 hover:bg-stone-600"
        type="button"
        @click="postPeriod"
      >
        Post
      </button>
    </div>
  </div>

  <!-- update form -->
  <div class="relative m-4 rounded bg-stone-50" v-else>
    <img
      src="@/assets/close.png"
      class="absolute right-2 top-2 h-3 w-3"
      @click="deselectPeriod"
    />
    <div class="p-3 shadow-xl">
      <!-- <p class="font-mono text-lg font-semibold">
        Current id: {{ selectedPeriod.ID }}
      </p> -->
      <label for="type" class="font-mono text-lg font-semibold">Type</label
      ><br />
      <select
        name="type"
        id="type"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="selectedPeriod.Type"
      >
        <option v-for="t in types" :value="t">{{ t }}</option>
      </select>
      <label for="title" class="font-mono text-lg font-semibold">Title</label
      ><br />
      <input
        type="text"
        id="title"
        name="title"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="selectedPeriod.Title"
      />
      <label for="description" class="font-mono text-lg font-semibold"
        >Description</label
      ><br />
      <textarea
        id="description"
        name="description"
        rows="4"
        class="mb-1 w-full resize-none rounded bg-stone-300 p-1 focus:outline-stone-700"
        v-model="selectedPeriod.Description"
      ></textarea>
      <label for="duration" class="font-mono text-lg font-semibold"
        >Duration</label
      ><br />
      <input
        type="text"
        id="duration"
        name="duration"
        class="mb-3 w-full rounded bg-stone-300 p-1 focus:outline-stone-700"
        :value="durationFormula"
        @input="updateDuration"
      />
      <button
        class="mb-1 w-full justify-center rounded bg-stone-700 text-center font-mono text-xl font-bold text-stone-50 hover:bg-stone-600"
        type="button"
        @click="updatePeriod"
      >
        Update
      </button>
      <button
        class="w-full justify-center rounded bg-stone-800 text-center font-mono text-xl font-bold text-stone-50 hover:bg-stone-900"
        type="button"
        @click="deletePeriod"
      >
        Delete
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { Period } from "@/components/API_Period.js";
import Mexp from "math-expression-evaluator";
let mexp = new Mexp();

export default {
  props: {
    types: { type: Array<string>, required: true },
    newPeriod: { type: Period, required: true },
    selectedPeriod: { type: Period, required: true },
  },
  data() {
    return {
      durationFormula: "" as string,
    };
  },
  watch: {
    // watches if a period is selected, to set the initial text for the duration amount
    selectedPeriod() {
      if (this.selectedPeriod.Date_Stamp == 0) this.durationFormula = "";
      else this.durationFormula = String(this.selectedPeriod.Duration);
    },
  },
  emits: ["postPeriod", "updatePeriod", "deletePeriod", "deselectPeriod"],
  methods: {
    // emit methods here, since they are easier to find like that
    postPeriod() {
      this.$emit("postPeriod");
    },
    deletePeriod() {
      this.$emit("deletePeriod");
    },
    updatePeriod() {
      this.$emit("updatePeriod");
    },
    deselectPeriod() {
      this.$emit("deselectPeriod");
    },

    // >> evaluates the durationFormula
    // & sets its value to either the new or selected period
    updateDuration(equation: Event) {
      try {
        if (equation.target == null) {
          return;
        }
        let target = equation.target as HTMLInputElement;
        let val = mexp.eval(target.value, [], []);
        if (this.selectedPeriod.Date_Stamp == 0) {
          this.newPeriod.Duration = val;
        } else {
          this.selectedPeriod.Duration = val;
        }
      } catch {
        // if an error occours then that simply means the formula wasn't correct
        // but it's not an issue since it will trigger while typing it out
      }
    },
  },
};
</script>
