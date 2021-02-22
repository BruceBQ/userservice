db.pages.drop();
db.pages.insertMany([
  // Bản đồ
  {
    name: "sitemap",
    displayName: "Bản đồ",
    path: "/admin/sitemap",
    description: "",
    isSettings: false,
    permissions: {
      get_parking: { compulsory: true },
      get_no_parking: { compulsory: true },
      get_camera: { compulsory: true },
      get_single_stream: {},
      edit_stream: {},
    },
  },

  // Danh sách theo dõi
  {
    name: "followlist",
    displayName: "Danh sách theo dõi",
    path: "/admin/followlist",
    description: "",
    isSettings: false,
    permissions: {
      get_stream: { compulsory: true },
      edit_stream: {},
      delete_stream: {},
    },
  },

  // Vi phạm
  {
    name: "violations",
    displayName: "Vi phạm",
    path: "/admin/violations",
    description: "",
    isSettings: false,
    permissions: {
      get_violations: { compulsory: true },
      statistical_violations: { compulsory: true },
      approve_violation: {},
      upapprove_violation: {},
      edit_violation: {},
      delete_violations: {},
      export_violation_report: {},
      export_statistical_violations_report: {},
    },
  },

  // Video
  {
    name: "video",
    displayName: "Quản lý video",
    path: "/admin/videos",
    description: "",
    isSettings: false,
    permissions: {
      get_videos: { compulsory: true },
      delete_videos: {},
    },
  },

  // Máy tính phân tán
  {
    name: "local_computers",
    displayName: "Quản lý máy tính phân tán",
    path: "/admin/local_computers",
    description: "",
    isSettings: false,
    permissions: {},
  },

  // Quản lý thu phí
  {
    name: "manage_fee",
    displayName: "Quản lý thu phí",
    path: "/admin/manage_fee",
    description: "",
    isSettings: false,
    permissions: {},
  },

  // Tin tức
  {
    name: "posts",
    displayName: "Thông báo, tin tức",
    path: "/admin/posts",
    description: "",
    isSettings: false,
    permissions: {
      get_posts: { compulsory: true },
      get_single_post: {},
      create_post: {},
      update_post: {},
      delete_post: {},
    },
  },

  // Quản lý camera
  {
    name: "manage_camera",
    displayName: "Quản lý camera",
    path: "/admin/manage_camera",
    description: "",
    isSettings: true,
    permissions: {
      get_camera: { compulsory: true },
      get_single_camera: {},
      add_camera: {},
      edit_camera: {},
      delete_camera: {},
    },
  },

  // Quản lý bãi đỗ
  {
    name: "manage_parking",
    displayName: "Quản lý bãi đỗ",
    path: "/admin/manage_camera",
    description: "",
    isSettings: true,
    permissions: {
      get_parking: { compulsory: true },
      create_parking: {},
      update_parking: {},
      delete_parking: {},
      export_parking_report: {},
    },
  },

  // Quản lý cấm đỗ
  {
    name: "manage_no_parking",
    displayName: "Quản lý khu vực cấm đố",
    path: "/admin/manage_no_parking",
    description: "",
    isSettings: true,
    permissions: {
      get_no_parking: { compulsory: true },
      get_single_no_parking: {},
      create_no_parking: {},
      update_no_parking: {},
      delete_no_parking: {},
      export_no_parking_report: {},
    },
  },

  // Quản lý người dùng
  {
    name: "manage_users",
    displayName: "Quản lý người dùng",
    path: "/admin/manage_users",
    description: "",
    isSettings: true,
    permissions: {
      get_users: { compulsory: true },
      create_user: {},
      get_single_user: {},
      update_user: {},
      delete_user: {},
    },
  },

  // Quản lý nhóm người dùng
  {
    name: "manage_roles",
    displayName: "Quản lý nhóm người dùng",
    path: "/admin/manage_roles",
    description: "",
    isSettings: true,
    permissions: {
      get_roles: { compulsory: true },
      get_single_role: {},
      create_role: {},
      update_role: {},
      delete_role: {},
      //   get_apis: { compulsory: true },
      //   get_pages: { compulsory: true },
    },
  },

  // Cài đặt hệ thống
  {
    name: "settings",
    displayName: "Cài đặt hệ thống",
    path: "/admin/settings",
    description: "",
    isSettings: true,
    permissions: {},
  },
]);
