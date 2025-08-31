local common = require("common")

local M = {}

M.load=function()	
	local filename = sys.get_save_file(common.app_name, "highscore")
	local data = sys.load(filename) 
	return data.highscore or 0
end

M.save = function(high_score)
	local filename = sys.get_save_file(common.app_name, "highscore")
	sys.save(filename, { highscore = high_score })  
end

return M