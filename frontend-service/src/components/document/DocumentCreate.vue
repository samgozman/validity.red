<template>
  <form @submit.prevent="">
    <div
      class="grid grid-cols-1 gap-6 py-5 lg:px-10 lg:grid-cols-3 lg:bg-base-200 rounded-box"
    >
      <div
        class="card col-span-1 row-span-1 shadow-lg lg:col-span-2 bg-base-100"
      >
        <div class="card-body form-control">
          <input
            type="text"
            placeholder="Document title"
            class="input w-full"
          />
          <label class="label">
            <span class="label-text text-base-300 px-3">
              Avoid unnecessary details.
            </span>
          </label>
        </div>
      </div>

      <div
        class="card col-span-1 row-span-1 shadow-lg lg:col-span-1 bg-base-100"
      >
        <div class="card-body">
          <select class="select w-full px-0">
            <option disabled selected>Select document type</option>
            <option>Type 1</option>
            <option>Type 2</option>
            <option>Type 3</option>
          </select>
        </div>
      </div>

      <div
        class="card col-span-1 row-span-2 shadow-lg lg:col-span-2 bg-base-100"
      >
        <div class="card-body form-control">
          <textarea
            class="textarea h-full"
            placeholder="Description (optional)"
          ></textarea>
          <label class="label">
            <span class="label-text text-base-300 px-3">
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
        <div class="card-body form-control">
          <label class="input-group">
            <input type="date" class="input input-bordered w-full" />
            <span>Expiration</span>
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
              <input type="checkbox" class="toggle toggle-primary" checked />
            </label>
          </div>
          <button class="btn btn-primary" type="submit">Save</button>
        </div>
      </div>
    </div>
  </form>
</template>
