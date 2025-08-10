components {
  id: "sound_manager"
  component: "/sound_manager.script"
}
components {
  id: "peg_hit_1"
  component: "/assets/sounds/pickupCoin.wav"
}
components {
  id: "peg_hit_2"
  component: "/assets/sounds/Pickup10.wav"
}
components {
  id: "peg_hit_3"
  component: "/assets/sounds/pickupCoin.wav"
}
components {
  id: "peg_hit_4"
  component: "/assets/sounds/PowerUp3.wav"
}
components {
  id: "peg_hit_5"
  component: "/assets/sounds/PowerUp4.wav"
}
components {
  id: "peg_hit_6"
  component: "/assets/sounds/PowerUp5.wav"
}
embedded_components {
  id: "peg_hit"
  type: "sound"
  data: "sound: \"/assets/sounds/peg_hit.wav\"\n"
  ""
}
