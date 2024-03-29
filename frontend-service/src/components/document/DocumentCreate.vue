<template>
  <form @submit.prevent="submit">
    <div
      class="grid grid-cols-1 gap-6 py-5 lg:px-10 lg:grid-cols-3 lg:bg-base-200 rounded-box"
    >
      <div
        class="card-card-compact card xl:card-normal col-span-1 row-span-1 shadow-lg lg:col-span-2 bg-base-100"
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
            <span class="label-text-alt text-base-300">
              {{ title.length }}/100
            </span>
          </label>
        </div>
      </div>

      <div
        class="card-card-compact card xl:card-normal col-span-1 row-span-2 shadow-lg lg:col-span-1 bg-base-100"
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
              :min="minDate"
              class="input input-bordered w-[100vw]"
              v-model="expiresAt"
            />
            <span>Expiration</span>
          </label>
        </div>
      </div>

      <div
        class="card-card-compact card xl:card-normal col-span-1 row-span-2 shadow-lg lg:col-span-2 bg-base-100"
      >
        <div class="card-body form-control">
          <h2 class="card-title">Description</h2>
          <textarea
            class="textarea h-full min-h-[20vh]"
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
            <span class="label-text-alt text-base-300">
              {{ description.length }}/500
            </span>
          </label>
        </div>
      </div>
      <div
        class="card-card-compact card xl:card-normal col-span-1 row-span-1 shadow-lg lg:col-span-1 bg-base-100"
      >
        <div class="card-body">
          <div v-if="!isEditMode" class="form-control">
            <label class="label cursor-pointer">
              <span class="label-text">Add default notification?</span>
              <input type="checkbox" class="toggle toggle-primary"
              v-model="createDefaultNotification" checked:="true" />
            </label>
          </div>
          <button class="btn btn-primary" type="submit">Save</button>
          <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
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
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      title: "",
      description: "",
      type: "Select document type",
      minDate: new Date().toISOString().split("T")[0],
      expiresAt: "",
      createDefaultNotification: true,
      typeOptions: DocumentType.types,
      documentTypeId: 0,
      isEditMode: false,
      documentId: "",
      errorMsg: "",
    };
  },
  methods: {
    isEmpty(str: string, msg: string): boolean {
      if (str === "") {
        this.errorMsg = msg;
        return true;
      }
      return false;
    },
    isLonger(str: string, len: number, msg: string) {
      if (str.length > len) {
        this.errorMsg = msg;
        return true;
      }
      return false;
    },
    isValidDate(dateString: string) {
      const parsed = Date.parse(dateString);
      if (isNaN(parsed)) {
        this.errorMsg = "Invalid date. Please use the format YYYY-MM-DD";
        return true;
      }
      return false;
    },
    isExpired(date: Date) {
      if (date < new Date()) {
        this.errorMsg = "Expiration date is in the past!";
        return false;
      }
    },
    async createDocument(expirationDate: Date) {
      this.documentId = await DocumentService.createOne({
        title: this.title,
        description: this.description,
        type: this.documentTypeId,
        expiresAt: expirationDate,
      });
    },
    async updateDocument(expirationDate: Date) {
      await DocumentService.updateOne({
        id: this.documentId,
        title: this.title,
        description: this.description,
        type: this.documentTypeId,
        expiresAt: expirationDate,
      });
    },
    async submit() {
      try {
        // 0. Clear the error message
        this.errorMsg = "";

        // 1. Validate form inputs
        if (this.isEmpty(this.title, "Title is required!")) return;
        if (this.isLonger(this.title, 100, "Title is too long!")) return;
        if (this.isLonger(this.description, 500, "Description is too long!"))
          return;
        if (this.isEmpty(this.expiresAt, "Expiration date is required!"))
          return;
        if (this.isValidDate(this.expiresAt)) return;

        const expirationDate = new Date(Date.parse(this.expiresAt));
        const typeIndex = [...this.typeOptions.values()]
          .map((d) => d.name)
          .indexOf(this.type);
        this.documentTypeId = typeIndex === -1 ? 0 : typeIndex;

        if (!this.isEditMode) {
          if (this.isExpired(expirationDate)) return;

          // 2. Create document and get its id
          await this.createDocument(expirationDate);
          // 3. Create default notification if needed
          if (this.createDefaultNotification) {
            expirationDate.setHours(10);
            await NotificationService.createOne({
              date: expirationDate,
              documentId: this.documentId,
            });
          }
        } else {
          // If it is the edit mode, we only need to update the document
          await this.updateDocument(expirationDate);
        }

        // 4. Redirect to document page
        this.$router.push({
          name: "document",
          params: { id: this.documentId },
        });
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
  async beforeMount() {
    if (this.$route.params.id) {
      this.createDefaultNotification = false;
      this.isEditMode = true;
      this.documentId =
        typeof this.$route.params.id === "string" ? this.$route.params.id : "";
    }

    if (this.isEditMode) {
      try {
        const document = await DocumentService.getOne(this.documentId);
        this.title = document.title || "";
        this.description = document.description || "";
        this.type = DocumentType.getName(document.type);
        this.expiresAt = new Date(document.expiresAt)
          .toISOString()
          .split("T")[0];
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error, this.$router);
      }
    }
  },
});
</script>
