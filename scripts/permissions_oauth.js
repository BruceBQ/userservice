db.permissions_oauth.drop();

db.permissions_oauth.insertMany([
  {
    name: "get_plate_numbers",
    displayName: "Lấy danh sách biển số của người dùng",
    method: "GET",
    path: "^/social/[a-f0-9]{24}/plate_numbers",
    description: "",
  },
  {
    name: "add_plate_number",
    displayName: "Thêm biển số vào danh sách của người dùng",
    method: "POST",
    path: "^/social/[a-f0-9]{24}/plate_numbers",
    description: "",
  },
  {
    name: "delete_plate_number",
    displayName: "Xóa biển số ra khỏi danh sách của người dùng",
    method: "DELETE",
    path: "^/social/[a-f0-9]{24}/plate_numbers",
    description: "",
  },
  {
    name: "update_status",
    displayName: "Cập nhật trạng thái người dùng",
    method: "PUT",
    path: "^/social/status$",
    description: "",
  },
  {
    name: "logout",
    displayName: "Đăng xuất",
    method: "POST",
    path: "^/social/logout$",
    description: "",
  },
]);
