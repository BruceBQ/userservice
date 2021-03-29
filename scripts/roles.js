db.roles.drop();
db.roles.insertMany([
  {
    name: "Super Admin",
    description: "",
    pages: {
      sitemap: { enabled: true },
      followlist: { enabled: true },
      violations: { enabled: true },
      local_computers: { enabled: true },
      manage_fee: { enabled: true },
      posts: { enabled: true },
      manage_camera: { enabled: true },
      manage_parking: { enabled: true },
      manage_no_parking: { enabled: true },
      manage_users: { enabled: true },
      manage_roles: { enabled: true },
      record_videos: { enabled: true },
      settings: { enabled: true },
    },
    permissions: {
      //parking
      get_parking: { enabled: true },
      create_parking: { enabled: true },
      update_parking: { enabled: true },
      get_single_parking: { enabled: true },
      delete_parking: { enabled: true },
      export_parking_report: { enabled: true },
      // no_parking
      get_no_parking: { enabled: true },
      create_no_parking: { enabled: true },
      update_no_parking: { enabled: true },
      get_single_no_parking: { enabled: true },
      delete_no_parking: { enabled: true },
      export_no_parking_report: { enabled: true },
      // posts
      get_posts: { enabled: true },
      create_post: { enabled: true },
      update_post: { enabled: true },
      get_single_post: { enabled: true },
      delete_post: { enabled: true },
      // violations
      get_violations: { enabled: true },
      approve_violation: { enabled: true },
      unapprove_violation: { enabled: true },
      edit_violation: { enabled: true },
      export_violation_report: { enabled: true },
      delete_violations: { enabled: true },
      statistical_violations: { enabled: true },
      export_statistical_violations_report: { enabled: true },
      get_video_violation: { enabled: true },
      // Camera
      get_camera: { enabled: true },
      get_single_camera: { enabled: true },
      get_camera_group: { enabled: true },
      get_camera_zone: { enabled: true },
      add_camera: { enabled: true },
      edit_camera: { enabled: true },
      delete_camera: { enabled: true },

      // Stream
      get_stream: { enabled: true },
      get_single_stream: { enabled: true },
      edit_stream: { enabled: true },
      delete_stream: { enabled: true },

      // Roles
      get_roles: { enabled: true },
      get_single_role: { enabled: true },
      create_role: { enabled: true },
      update_role: { enabled: true },
      delete_role: { enabled: true },
      // Users
      get_users: { enabled: true },
      create_user: { enabled: true },
      get_single_user: { enabled: true },
      update_user: { enabled: true },
      delete_user: { enabled: true },
      // Videos
      get_videos: { enabled: true },
      delete_videos: { enabled: true },
    },
    builtin: true,
  },
]);
