const a = {
  _id: ObjectId("5fd833b11e576d46f80a7d29"),
};

const b = {
  role_object_id: { $toObjectId: "$role_id" },
};
const c = {
  from: "roles",
  let: { role_object_id: "$role_object_id" },
  pipeline: [
    {
      $match: {
        $expr: {
          $and: [{ $eq: ["$_id", "$$role_object_id"] }],
        },
      },
    },
  ],
  as: "roles",
};

const d = {
  from: "roles",
  localField: "role_object_id",
  foreignField: "_id",
  as: "string",
};
