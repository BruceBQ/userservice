db.roles.drop();
db.roles.insertMany([
  {
    name: "Super Admin",
    description: "",
    pages: {
      sitemap: {
        enabled: true,
      },
      followlist: {
        enabled: true,
      },
      violations: {
        enabled: true,
      },
      local_computers: {
        enabled: true,
      },
      manage_fee: {
        enabled: true,
      },
      posts: {
        enabled: true,
      },
      manage_camera: {
        enabled: true,
      },
      manage_parking: {
        enabled: true,
      },
      manage_no_parking: {
        enabled: true,
      },
      manage_users: {
        enabled: true,
      },
      manage_roles: {
        enabled: true,
      },
      settings: {
        enabled: true,
      },
    },
    permissions: {
      get_political: {
        enabled: true,
      },
      get_parking: {
        enabled: true,
      },
      create_parking: {
        enabled: true,
      },
      update_parking: {
        enabled: true,
      },
      get_single_parking: {
        enabled: true,
      },
      delete_parking: {
        enabled: true,
      },
      get_no_parking: {
        enabled: true,
      },
      create_no_parking: {
        enabled: true,
      },
      update_no_parking: {
        enabled: true,
      },
      get_single_no_parking: {
        enabled: true,
      },
      delete_no_parking: {
        enabled: true,
      },
      get_posts: {
        enabled: true,
      },
      create_post: {
        enabled: true,
      },
      update_post: {
        enabled: true,
      },
      get_single_post: {
        enabled: true,
      },
      delete_post: {
        enabled: true,
      },
      get_violations: {
        enabled: true,
      },
      get_single_violation: {
        enabled: true,
      },
      update_violation: {
        enabled: true,
      },
      delete_violation: {
        enabled: true,
      },
      get_cameras: {
        enabled: true,
      },
      get_single_camera: {
        enabled: true,
      },
      delete_camera: {
        enabled: true,
      },
      get_camera_group: {
        enabled: true,
      },
      get_zones: {
        enabled: true,
      },
      create_camera: {
        enabled: true,
      },
      update_camera: {
        enabled: true,
      },
      get_roles: {
        enabled: true,
      },
      get_single_role: {
        enabled: true,
      },
      create_role: {
        enabled: true,
      },
      update_role: {
        enabled: true,
      },
      delete_role: {
        enabled: true,
      },
      get_permissions: {
        enabled: true,
      },
      get_accounts: {
        enabled: true,
      },
      get_single_account: {
        enabled: true,
      },
      create_account: {
        enabled: true,
      },
      update_account: {
        enabled: true,
      },
      delete_account: {
        enabled: true,
      },
      get_user: {
        enabled: true,
      },
      update_user: {
        enabled: true,
      },
      get_followlist: {
        enabled: true,
      },
      get_camera_unfollowed: {
        enabled: true,
      },
      add_camera_to_followlist: {
        enabled: true,
      },
      remove_camera_from_followlist: {
        enabled: true,
      },
      fee: {
        enabled: true,
      },
      get_notifications: {
        enabled: true,
      },
      mark_read: {
        enabled: true,
      },
    },
    builtin: true,
  },
  {
    name: "Admin",
    description: "",
    pages: ["sitemap", "followlist"],
    api: ["get_notifications", "mark_read"],
  },
]);
