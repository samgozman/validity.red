<template>
  <form @submit.prevent="submit">
    <div
      class="grid grid-cols-1 gap-6 py-5 lg:px-10 lg:grid-cols-3 lg:bg-base-200 rounded-box"
    >
      <div
        class="card col-span-1 row-span-1 shadow-lg lg:col-span-2 bg-base-100"
      >
        <div class="card-body form-control">
          <h2 class="card-title">Title</h2>
          <input
            type="text"
            placeholder="Document title"
            minlength="3"
            maxlength="100"
            class="input w-full"
            v-model="title"
          />
          <label class="label">
            <span class="label-text text-base-300">
              Avoid unnecessary details.
            </span>
          </label>
        </div>
      </div>

      <div
        class="card col-span-1 row-span-2 shadow-lg lg:col-span-1 bg-base-100"
      >
        <div class="card-body">
          <h2 class="card-title">Document settings</h2>
          <select class="select select-bordered w-full" v-model="type">
            <option disabled selected>Select document type</option>
            <option
              v-for="(option, index) in typeOptions.values()"
              v-bind:key="index"
            >
              {{ option.name }}
            </option>
          </select>
          <div class="divider">Expires at</div>
          <label class="input-group">
            <input
              type="date"
              class="input input-bordered w-full"
              v-model="expiresAt"
            />
            <span>Expiration</span>
          </label>
        </div>
      </div>

      <div
        class="card col-span-1 row-span-2 shadow-lg lg:col-span-2 bg-base-100"
      >
        <div class="card-body form-control">
          <h2 class="card-title">Description</h2>
          <textarea
            class="textarea h-full"
            placeholder="Description (optional)"
            maxlength="500"
            v-model="description"
          ></textarea>
          <label class="label">
            <span class="label-text text-base-300">
              Avoid unnecessary details. We don't need exact information about
              your document. <br />
              Please use this field only to distinguish documents from each
              other.
            </span>
          </label>
        </div>
      </div>
      <div
        class="card col-span-1 row-span-1 shadow-lg lg:col-span-1 bg-base-100"
      >
        <div class="card-body">
          <div class="form-control">
            <label class="label cursor-pointer">
              <span class="label-text">Add default notification?</span>
              <input
                type="checkbox"
                class="toggle toggle-primary"
                v-model="createDefaultNotification"
                checked
              />
            </label>
          </div>
          <button class="btn btn-primary" type="submit">Save</button>
          <div v-show="error" class="badge badge-error badge-outline w-full">
            {{ errorMsg }}
          </div>
        </div>
      </div>
    </div>
  </form>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { DocumentService } from "./DocumentService";
import { NotificationService } from "./NotificationService";
import { DocumentType } from "./DocumentType";

export default defineComponent({
  data() {
    return {
      title: "",
      description: "",
      type: "Select document type",
      expiresAt: "",
      createDefaultNotification: true,
      typeOptions: DocumentType.types,
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    isEmpty(str: string, msg: string): boolean {
      if (str === "") {
        this.error = true;
        this.errorMsg = msg;
        return true;
      }
      return false;
    },
    isLonger(str: string, len: number, msg: string) {
      if (str.length > len) {
        this.error = true;
        this.errorMsg = msg;
        return true;
      }
      return false;
    },
    isExpired(dstr: string) {
      const parsed = Date.parse(dstr);
      if (isNaN(parsed)) {
        this.error = true;
        this.errorMsg = "Invalid date. Please use the format YYYY-MM-DD";
        return true;
      }
      if (new Date(parsed) < new Date()) {
        this.error = true;
        this.errorMsg = "Expiration date is in the past!";
        return false;
      }
    },
    async submit() {
      try {
        // 0. Clear the error message
        this.error = false;
        this.errorMsg = "";

        // 1. Validate form inputs
        if (this.isEmpty(this.title, "Title is required!")) return;
        if (this.isLonger(this.title, 100, "Title is too long!")) return;
        if (this.isLonger(this.description, 500, "Description is too long!"))
          return;
        if (this.isLonger(this.title, 100, "Title is too long!")) return;
        if (this.isEmpty(this.expiresAt, "Expiration date is required!"))
          return;
        if (this.isExpired(this.expiresAt)) return;

        // 2. Create document and get its id
        const expirationDate = new Date(Date.parse(this.expiresAt));
        const typeIndex = [...this.typeOptions.values()]
          .map((d) => d.name)
          .indexOf(this.type);
        const documentType = typeIndex === -1 ? 0 : typeIndex;
        const documentId = await DocumentService.createOne({
          title: this.title,
          description: this.description,
          type: documentType,
          expiresAt: expirationDate,
        });

        // 3. Create default notification if needed
        if (this.createDefaultNotification) {
          await NotificationService.createOne({
            date: expirationDate,
            documentId: documentId,
          });
        }

        // 4. Redirect to document page
        this.$router.push({
          name: "document",
          params: { id: documentId },
        });
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>
