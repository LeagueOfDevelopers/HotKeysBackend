using Microsoft.AspNetCore.Mvc;
using System;
using BisinessLogic;
using System.Collections.Generic;

namespace HotKeys.Controllers
{
    public class ProgramController : ControllerBase
    {
        [HttpGet]
        [Route("programs/{id}")]
        public Program GetProgramByID(Guid id)
        {
            return new Program(id, "MaratPidor", new List<HotKey>());
        }
    }
}