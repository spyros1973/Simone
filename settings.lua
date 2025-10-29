local M = {}

M.sound=true
--M.scanlines=false
M.volume=0.5
M.sound_theme="fx" -- notes|fx

M.load=function()
	local filename = sys.get_save_file("sys_save_load", "settings")
	print("settings: "..filename)
	if not sys.exists(filename) then return	end
	local data = sys.load(filename) 
	M.sound=data.sound or false
	--M.scanlines=data.scanlines or false
	M.volume=data.volume or 0.5
	M.sound_theme=data.sound_theme or "notes"
	return data
end

M.save = function()
	local filename = sys.get_save_file("sys_save_load", "settings")
	sys.save(filename, { sound = M.sound, volume = M.volume, sound_theme=M.sound_theme}) --, scanlines = M.scanlines, volume = M.volume })  
end

return M