components {
  id: "screen-overlay"
  component: "/game/screen-overlay.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"black_circle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/main.atlas\"\n"
  "}\n"
  ""
}
