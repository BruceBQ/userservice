db.permissions.drop();

db.permissions.insertMany([
  // Parking
  { name: "get_parking", displayName: "Xem danh sách bãi đỗ xe", description: "" },
  { name: "create_parking", displayName: "Thêm bãi đỗ xe", description: "" },
  { name: "update_parking", displayName: "Sửa thông tin bãi đỗ xe", description: "" },
  { name: "get_single_parking", displayName: "Xem thông tin bãi đỗ", description: "" },
  { name: "delete_parking", displayName: "Xóa bãi đỗ xe", description: "" },
  { name: "export_parking_report", displayName: "Xuất báo cáo danh sách bãi đỗ", description: "" },
//   { name: "update_empty_slots", displayName: "Cập nhật số vị trí trống của bãi đỗ", descriptions: "" },

  // No parking
  { name: "get_no_parking", displayName: "Xem danh sách khu vực cấm đỗ", description: "" },
  { name: "create_no_parking", displayName: "Thêm khu vực cấm đỗ", description: "" },
  { name: "update_no_parking", displayName: "Sửa thông tin khu vực cấm đỗ", description: "" },
  { name: "get_single_no_parking", displayName: "Xem thông tin khu vực cấm đỗ", description: "" },
  { name: "delete_no_parking", displayName: "Xóa khu vực cấm đỗ", description: "" },
  { name: "export_no_parking_report", displayName: "Xuất báo cáo danh sách khu vực cấm đỗ", description: "" },
  // Posts
  { name: "get_posts", displayName: "Xem danh sách tin tức", description: "" },
  { name: "create_post", displayName: "Thêm tin tức", description: "" },
  { name: "update_post", displayName: "Cập nhật nội dung tin tức", description: "" },
  { name: "get_single_post", displayName: "Xem thông tin chi tiết tin tức", description: "" },
  { name: "delete_post", displayName: "Xóa tin tức", description: "" },

  // Violations
  { name: "get_vehicletype", displayName: "Lấy danh sách loại phương tiện" },
  { name: "upload_file", displayName: "Upload ảnh/video vi phạm" },
  { name: "create_violation", displayName: "Thêm vi phạm" },
  { name: "get_violations", displayName: "Lấy danh sách vi phạm", description: "" },
  { name: "approve_violation", displayName: "Duyệt vi phạm", description: "" },
  { name: "upapprove_violation", displayName: "Bỏ duyệt vi phạm vi phạm", description: "" },
  { name: "edit_violation", displayName: "Chỉnh sửa thông tin vi phạm", description: "" },
  { name: "delete_violations", displayName: "Xóa vi phạm", description: "" },
  { name: "export_violation_report", displayName: "Xuất biên bản vi phạm", description: "" },
  { name: "statistical_violations", displayName: "Thống kê vi phạm", description: "" },
  { name: "export_statistical_violations_report", displayName: "Xuất báo cáo thống kê vi phạm", description: "" },
  { name: "get_video_violation", displayName: "Lấy video vi phạm", description: "" },

  // Videos
  { name: "get_videos", displayName: "Lấy danh sách video", description: "" },
  { name: "delete_videos", displayName: "Xóa video", description: "" },

  // Cameras
  { name: "get_camera", displayName: "Xem danh sách camera", description: "" },
  { name: "get_single_camera", displayName: "Xem thông tin camera", description: "" },
  { name: "add_camera", displayName: "Thêm mới camera", description: "" },
  { name: "edit_camera", displayName: "Chỉnh sửa camera", description: "" },
  { name: "delete_camera", displayName: "Xóa camera", description: "" },
  { name: "get_camera_group", displayName: "Lấy danh sách nhóm camera", description: "" },
  { name: "get_camera_zone", displayName: "Lấy danh sách các vùng nhận dạng", description: "" },
  { name: "ptz_control", displayName: "Điểu khiển PTZ" },

  // Stream
  { name: "get_stream", displayName: "Lấy thông tin stream một danh sách camera", description: "" },
  { name: "get_single_stream", displayName: "Lấy thông tin stream một camera", description: "" },
  { name: "edit_stream", displayName: "Chỉnh sửa thông tin stream camera", description: "" },
  { name: "delete_stream", displayName: "Xóa stream camera", description: "" },

  // Roles
  { name: "get_roles", displayName: "Lấy danh sách nhóm người dùng", description: "" },
  { name: "get_single_role", displayName: "Lấy thông tin nhóm người dùng", description: "" },
  { name: "create_role", displayName: "Thêm nhóm người dùng", description: "" },
  { name: "update_role", displayName: "Sửa thông tin nhóm người dùng", description: "" },
  { name: "delete_role", displayName: "Xóa nhóm người dùng", description: "" },

  // Users
  { name: "get_users", displayName: "Xem danh sách người dùng", description: "" },
  { name: "create_user", displayName: "Thêm mới người dùng", description: "" },
  { name: "get_single_user", displayName: "Xem thông tin người dùng", description: "" },
  { name: "update_user", displayName: "Sửa thông tin người dùng", description: "" },
  { name: "delete_user", displayName: "Xóa người dùng", description: "" },

  

  // Parking Fee ???
//   { name: "fee", displayName: "Quản lý thu phí", description: "" },

  // Notifications
//   { name: "get_notifications", displayName: "Xem danh sách thông báo", description: "" },
//   { name: "mark_read", displayName: "Đánh dấu thông báo là đã đọc", description: "" },

  // Local Computer

  // Logs

  // Cài đặt
]);
