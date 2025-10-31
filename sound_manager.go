components {
  id: "sound_manager"
  component: "/sound_manager.script"
}
components {
  id: "peg_hit_1"
  component: "/assets/sounds/snd-1.wav"
}
components {
  id: "peg_hit_2"
  component: "/assets/sounds/snd-2.wav"
}
components {
  id: "peg_hit_3"
  component: "/assets/sounds/snd-3.wav"
}
components {
  id: "peg_hit_4"
  component: "/assets/sounds/snd-4.wav"
}
components {
  id: "error"
  component: "/assets/sounds/error.wav"
}
components {
  id: "powerUp"
  component: "/assets/sounds/powerup.wav"
}
components {
  id: "game_over"
  component: "/assets/sounds/game_over.wav"
}
embedded_components {
  id: "peg_hit"
  type: "sound"
  data: "sound: \"/assets/sounds/peg_hit.wav\"\n"
  ""
}
embedded_components {
  id: "do"
  type: "sound"
  data: "sound: \"/assets/sounds/do.wav\"\n"
  ""
}
embedded_components {
  id: "re"
  type: "sound"
  data: "sound: \"/assets/sounds/re.wav\"\n"
  ""
}
embedded_components {
  id: "mi"
  type: "sound"
  data: "sound: \"/assets/sounds/mi.wav\"\n"
  ""
}
embedded_components {
  id: "fa"
  type: "sound"
  data: "sound: \"/assets/sounds/fa.wav\"\n"
  ""
}
embedded_components {
  id: "tick"
  type: "sound"
  data: "sound: \"/assets/sounds/tick.wav\"\n"
  ""
}
embedded_components {
  id: "click"
  type: "sound"
  data: "sound: \"/assets/sounds/click.wav\"\n"
  ""
}
