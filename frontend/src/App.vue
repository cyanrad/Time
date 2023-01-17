<template>
  <div class="fixed top-0 left-0 flex h-full w-full flex-row">
    <div class="flex basis-1/3 flex-col">
      <SummeryTable
        :periods="periods"
        :types="types"
        :get-type-color="getTypeColor"
      />
      <PeriodForm
        :types="types"
        :new-period="newPeriod"
        :selected-period="selectedPeriod"
        @post-period="postPeriod()"
        @delete-period="deletePeriod()"
        @update-period="updatePeriod()"
        @deselect-period="deselectPeriod()"
      />
      <!-- spacer -->
      <div class="h-full"></div>
      <!-- spacer -->
    </div>
    <TimeBar
      :periods="periods"
      :get-type-color="getTypeColor"
      @period-pressed="selectPeriod"
    />
    <div class="m-2 flex w-full">
      <div class="m-2 flex basis-1/3 flex-col">
        <div class="mb-2 flex basis-1/2 flex-col">
          <AchievementsPanel
            :new-acheivement="newAchievement"
            :achievements="achievements"
            @post-achievements="postAchievement()"
            @delete-achievements="(index) => deleteAchievement(index)"
          />
        </div>
        <div class="mt-2 basis-1/2"><StatisticsTable /></div>
      </div>
      <div class="m-2 flex grow basis-2/3 flex-col">
        <div class="mb-2 basis-1/2"><SummeryDescription /></div>
        <div class="mt-2 basis-1/2">
          <Calender
            :current-date-stamp="selectedDateStamp"
            :date-stamp-to-date="dateStampToDate"
            @change-date="changeDate"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import PeriodForm from "./components/PeriodForm.vue";
import TimeBar from "./components/TimeBar.vue";
import SummeryTable from "./components/SummeryTable.vue";
import StatisticsTable from "./components/StatisticsTable.vue";
import SummeryDescription from "./components/SummeryDescription.vue";
import Calender from "./components/Calender.vue";
import AchievementsPanel from "./components/AchievementsPanel.vue";

import * as API_period from "@/components/API_Period";
import * as API_Achievement from "@/components/API_Achievement";
import { Period } from "@/components/API_Period";
import { Achievement } from "./components/API_Achievement";

// default data until implemented on server
const typesDefault: string[] = [
  "Work",
  "Positive",
  "People",
  "Neutral",
  "University",
  "Rest",
  "Useless",
  "Mandatory",
  "Unknown",
];
const typesColorsDefault: Map<string, Array<string>> = new Map([
  ["Positive", ["#FFF5FF", "#E3ACF9", "#8D5BA2"]],
  ["Work", ["#ECFEEF", "#C8FFD4", "#62956F"]],
  ["Rest", ["#FBFAEE", "#F5F471", "#888D00"]],
  ["Neutral", ["#e7e5e4", "#a3a3a3", "#404040"]],
  ["Unknown", ["#e7e5e4", "#a3a3a3", "#404040"]],
  ["University", ["#F9F8FF", "#8A90FF", "#3E51B6"]],
  ["People", ["#FFF7EF", "#EE9757", "#9A5011"]],
  ["Useless", ["#FFF4F3", "#FD8A8A", "#A94045"]],
  ["Mandatory", ["#FFF4F3", "#FD8A8A", "#A94045"]],
  ["Useless", ["#e7e5e4", "#a3a3a3", "#404040"]],
]);

export default {
  data() {
    return {
      periods: new Array<Period>(),
      achievements: new Array<Achievement>(),

      types: typesDefault,
      typesColors: typesColorsDefault,

      newPeriod: this.getNewPeriod(), // the new period to be added when submit is perssed
      newAchievement: this.getNewAchievement(),
      selectedPeriodInitialData: this.getNewPeriod(),
      selectedPeriodIndex: null as number | null, // null if no period is selected
      selectedDateStamp: 0 as number,
    };
  },

  computed: {
    // >> gets the period selected from the TimeBar.vue
    // @ no period selected (selectedPeriodIndex == null) return empty Period
    selectedPeriod(): Period {
      return this.selectedPeriodIndex != null
        ? this.periods[this.selectedPeriodIndex]
        : this.getNewPeriod();
    },
  },

  methods: {
    // >> returns empty obj
    getNewPeriod(): Period {
      return new Period(0, 0, 0, 0, "", "", "");
    },
    getNewAchievement(): Achievement {
      return new Achievement(0, 0, "");
    },

    // >> period (post, update, delete) methods
    async postPeriod() {
      await API_period.postPeriods(
        this.periods,
        this.newPeriod,
        this.selectedDateStamp
      );
      this.newPeriod = this.getNewPeriod();
    },
    async deletePeriod() {
      if (this.selectedPeriodIndex != null) {
        // temp index exists cuz index needs to be null before deleting
        let tempIndex = this.selectedPeriodIndex;
        this.selectedPeriodIndex = null;
        await API_period.deletePeriods(this.periods, tempIndex);
      }
    },
    async updatePeriod() {
      if (this.selectedPeriodIndex != null) {
        await API_period.updatePeriods(
          this.periods,
          this.periods[this.selectedPeriodIndex],
          this.selectedPeriodIndex
        );
        this.selectedPeriodIndex = null;
      }
    },

    // >> achievement (post, delete) methods
    async postAchievement() {
      await API_Achievement.postAchievements(
        this.achievements,
        this.newAchievement,
        this.selectedDateStamp
      );
      this.newAchievement = this.getNewAchievement();
    },
    async deleteAchievement(index: number) {
      await API_Achievement.deleteAchievements(this.achievements, index);
    },

    deselectPeriod() {
      if (this.selectedPeriodIndex != null) {
        // resetting the period if it was modified in the update period form
        this.periods[this.selectedPeriodIndex] = this.selectedPeriodInitialData;
        this.selectedPeriodIndex = null;
      }
    },

    selectPeriod(index: number) {
      this.selectedPeriodIndex = index;
    },

    // >> when the user chooses a new date from the calendar
    async changeDate(date: Date) {
      this.deselectPeriod();

      this.selectedDateStamp = this.dateToDateStamp(date);

      // getting data on that day
      this.periods = await API_period.getPeriods(this.selectedDateStamp);
      this.achievements = await API_Achievement.getAchievements(
        this.selectedDateStamp
      );
    },

    // >> datestamp methods
    // datestamp is a number representation of current day (ex: 20230113)
    dateToDateStamp(day: Date): number {
      let dd: number = day.getDate();
      let mm: number = day.getMonth() + 1;
      let yyyy: number = day.getFullYear();
      return yyyy * 10000 + mm * 100 + dd;
    },
    dateStampToDate(dateStamp: number): Date {
      let strDS: string = String(dateStamp);
      return new Date(
        parseInt(strDS.substring(0, 4)),
        parseInt(strDS.substring(4, 6)) - 1,
        parseInt(strDS.substring(6))
      );
    },

    // >> returns a color based on the type and how dark it is (0 being lightest)
    getTypeColor(type: string, darkness: number): string {
      let item = this.typesColors.get(type);
      if (item != undefined && darkness < 4) {
        return item[darkness];
      }
      return "";
    },
  },

  async created() {
    // setting date stamp to today
    this.selectedDateStamp = this.dateToDateStamp(new Date());
    // getting initial data
    this.periods = await API_period.getPeriods(this.selectedDateStamp);
    this.achievements = await API_Achievement.getAchievements(
      this.selectedDateStamp
    );
  },
  updated() {
    if (this.selectedPeriodIndex != null) {
      // making a copy of the data of the selected period
      // before it is updated by user, to reset if update is cancled
      this.selectedPeriodInitialData = Object.assign(
        {},
        this.periods[this.selectedPeriodIndex]
      );
    }
  },

  components: {
    PeriodForm,
    TimeBar,
    SummeryTable,
    StatisticsTable,
    SummeryDescription,
    Calender,
    AchievementsPanel,
  },
};
</script>
