<template>
  <div class="flex h-full w-full flex-col rounded bg-stone-50">
    <div
      v-for="(a, i) in achievements"
      class="m-1 mr-2 ml-2 flex h-7 flex-row items-center rounded-lg bg-stone-300 pl-2 pr-2 font-semibold"
    >
      <p class="w-full">{{ a.Text }}</p>
      <img src="@/assets/edit.png" class="mr-1 h-4 w-4" />
      <img
        src="@/assets/delete.png"
        class="h-4 w-4"
        @click="deleteAchievements(i)"
      />
    </div>

    <div
      v-if="!isAddingAchievemtn"
      class="m-2 flex h-7 rounded-lg bg-stone-300 text-center text-lg font-semibold"
      @click="switchToAddingMode()"
    >
      <img src="@/assets/addAchievement.png" class="m-auto h-5 w-5" />
    </div>

    <div
      class="m-2 flex h-7 flex-row items-center rounded-lg bg-stone-300 p-1"
      v-else
    >
      <img
        src="@/assets/close.png"
        class="h-3 w-3"
        @click="closeAddingMode()"
      />
      <input
        type="text"
        class="mr-2 ml-2 w-full rounded bg-stone-300 focus:outline-none"
        ref="achievement_field"
        v-model="newAcheivement.Text"
      />
      <img
        src="@/assets/addAchievement.png"
        class="h-5 w-5"
        @click="postAchievements()"
      />
    </div>
  </div>
</template>

<script lang="ts">
import * as API from "./API_Achievement";

export default {
  props: {
    achievements: { type: Array<API.Achievement>, requred: true },
    newAcheivement: { type: API.Achievement, required: true },
  },
  data() {
    return {
      isAddingAchievemtn: false as boolean,
    };
  },
  emits: ["postAchievements", "deleteAchievements", "updateAchievements"],
  methods: {
    switchToAddingMode() {
      this.isAddingAchievemtn = true;
    },
    closeAddingMode() {
      this.isAddingAchievemtn = false;
      this.newAcheivement.Text = "";
    },
    postAchievements() {
      this.$emit("postAchievements");
      this.isAddingAchievemtn = false;
    },
    deleteAchievements(index: number) {
      this.$emit("deleteAchievements", index);
    },
    updateAchievements() {
      this.$emit("updateAchievements");
    },
  },
};
</script>
