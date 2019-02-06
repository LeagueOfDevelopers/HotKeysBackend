using Microsoft.AspNetCore.Mvc;
using System;
using HotKeysLibrary.Entities;
using System.Collections.Generic;

namespace HotKeysAPI.Controllers
{
    public class ProgramController : ControllerBase
    {
        [HttpGet]
        [Route("programs/{id}")]
        public HotKeysLibrary.Entities.Program GetProgramByID(Guid id)
        {
            return new HotKeysLibrary.Entities.Program(id, "Figma", new List<HotKey>());
        }
    }
}